package util

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gorilla/websocket"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"net/http"
	"net/url"
	"time"
)

func HttpGet(ctx context.Context, url string, header map[string]string, data g.Map, result interface{}, proxyURL ...string) error {

	logger.Infof(ctx, "HttpGet url: %s, header: %+v, data: %+v, proxyURL: %v", url, header, data, proxyURL)

	client := g.Client().Timeout(config.Cfg.Http.Timeout * time.Second)

	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(proxyURL) > 0 {
		client.SetProxy(proxyURL[0])
	} else if config.Cfg.Http.ProxyUrl != "" {
		client.SetProxy(config.Cfg.Http.ProxyUrl)
	}

	response, err := client.Get(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		if err = response.Close(); err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpGet url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		if err = gjson.Unmarshal(bytes, result); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpPostJson(ctx context.Context, url string, header map[string]string, data, result interface{}, proxyURL ...string) error {

	logger.Infof(ctx, "HttpPostJson url: %s, header: %+v, data: %+v, proxyURL: %v", url, header, data, proxyURL)

	client := g.Client().Timeout(config.Cfg.Http.Timeout * time.Second)

	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(proxyURL) > 0 {
		client.SetProxy(proxyURL[0])
	} else if config.Cfg.Http.ProxyUrl != "" {
		client.SetProxy(config.Cfg.Http.ProxyUrl)
	}

	response, err := client.ContentJson().Post(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		if err = response.Close(); err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpPostJson url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		if err = gjson.Unmarshal(bytes, result); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func HttpPost(ctx context.Context, url string, header map[string]string, data, result interface{}, proxyURL ...string) error {

	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, proxyURL: %v", url, header, data, proxyURL)

	client := g.Client().Timeout(config.Cfg.Http.Timeout * time.Second)

	if header != nil {
		client.SetHeaderMap(header)
	}

	if len(proxyURL) > 0 {
		client.SetProxy(proxyURL[0])
	} else if config.Cfg.Http.ProxyUrl != "" {
		client.SetProxy(config.Cfg.Http.ProxyUrl)
	}

	response, err := client.Post(ctx, url, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	defer func() {
		if err = response.Close(); err != nil {
			logger.Error(ctx, err)
		}
	}()

	bytes := response.ReadAll()
	logger.Infof(ctx, "HttpPost url: %s, header: %+v, data: %+v, response: %s", url, header, data, string(bytes))

	if bytes != nil && len(bytes) > 0 {
		if err = gjson.Unmarshal(bytes, result); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func WebSocketClientOnlyReceive(ctx context.Context, wsURL string, result chan []byte, proxyURL ...string) (*websocket.Conn, error) {

	logger.Infof(ctx, "WebSocketClientOnlyReceive wsURL: %s", wsURL)

	client := gclient.NewWebSocket()

	client.HandshakeTimeout = config.Cfg.Http.Timeout * time.Second // 设置超时时间
	//client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置

	// 设置代理
	if len(proxyURL) > 0 {
		if proxyUrl, err := url.Parse(proxyURL[0]); err != nil {
			logger.Error(ctx, err)
		} else {
			client.Proxy = http.ProxyURL(proxyUrl)
		}
	} else if config.Cfg.Http.ProxyUrl != "" {
		if proxyUrl, err := url.Parse(config.Cfg.Http.ProxyUrl); err != nil {
			logger.Error(ctx, err)
		} else {
			client.Proxy = http.ProxyURL(proxyUrl)
		}
	}

	conn, _, err := client.Dial(wsURL, nil)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	entry := gtimer.AddSingleton(ctx, 30*time.Second, func(ctx context.Context) {
		logger.Debugf(ctx, "WebSocketClientOnlyReceive wsURL: %s, ping...", wsURL)
		if err = conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
			logger.Error(ctx, err)
			return
		}
	})

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

		defer entry.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error(ctx, err)
				return
			}
			logger.Infof(ctx, "messageType: %d, message: %s", messageType, string(message))

			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
				result <- message
			}, nil)
		}
	}, nil)

	return conn, nil
}

func WebSocketClient(ctx context.Context, wsURL string, messageType int, message []byte, result chan []byte, proxyURL ...string) (*websocket.Conn, error) {

	logger.Infof(ctx, "WebSocketClient wsURL: %s", wsURL)

	client := gclient.NewWebSocket()

	client.HandshakeTimeout = config.Cfg.Http.Timeout * time.Second // 设置超时时间
	//client.TLSClientConfig = &tls.Config{}   // 设置 tls 配置

	// 设置代理
	if len(proxyURL) > 0 {
		if proxyUrl, err := url.Parse(proxyURL[0]); err != nil {
			logger.Error(ctx, err)
		} else {
			client.Proxy = http.ProxyURL(proxyUrl)
		}
	} else if config.Cfg.Http.ProxyUrl != "" {
		if proxyUrl, err := url.Parse(config.Cfg.Http.ProxyUrl); err != nil {
			logger.Error(ctx, err)
		} else {
			client.Proxy = http.ProxyURL(proxyUrl)
		}
	}

	conn, _, err := client.Dial(wsURL, nil)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = conn.WriteMessage(messageType, message); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error(ctx, err)
				return
			}
			logger.Infof(ctx, "messageType: %d, message: %s", messageType, string(message))

			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
				result <- message
			}, nil)
		}
	}, nil)

	return conn, nil
}

func HttpDownloadFile(ctx context.Context, fileURL string, proxyURL ...string) []byte {

	logger.Infof(ctx, "HttpDownloadFile fileURL: %s", fileURL)

	client := g.Client().Timeout(config.Cfg.Http.Timeout * time.Second)

	if len(proxyURL) > 0 {
		client.SetProxy(proxyURL[0])
	} else if config.Cfg.Http.ProxyUrl != "" {
		client.SetProxy(config.Cfg.Http.ProxyUrl)
	}

	return client.GetBytes(ctx, fileURL)
}
