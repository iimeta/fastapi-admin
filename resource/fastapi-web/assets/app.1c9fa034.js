import{c as t,A as e}from"./index.e0d6ab41.js";function r(p){return t.post("/api/v1/app/create",p)}function n(p){return t.post("/api/v1/app/page",p)}function s(){return t.get("/api/v1/app/list")}function u(p){return t.post("/api/v1/app/delete",p)}function o(p){return t.get("/api/v1/app/detail",{params:p,paramsSerializer:a=>e.stringify(a)})}function c(p){return t.post("/api/v1/app/update",p)}function f(p){return t.post("/api/v1/app/change/status",p)}function g(p){return t.post("/api/v1/app/create/key",p)}function A(p){return t.post("/api/v1/app/key/config",p)}export{n as a,f as b,g as c,A as d,r as e,c as f,s as g,o as q,u as s};
