import{x as t,y as a}from"./index.74baba8a.js";function r(p){return t.post("/api/v1/app/create",p)}function n(p){return t.post("/api/v1/app/page",p)}function s(){return t.get("/api/v1/app/list")}function u(p){return t.post("/api/v1/app/delete",p)}function o(p){return t.get("/api/v1/app/detail",{params:p,paramsSerializer:e=>a.stringify(e)})}function f(p){return t.post("/api/v1/app/update",p)}function c(p){return t.post("/api/v1/app/create/key",p)}function y(p){return t.post("/api/v1/app/key/config",p)}export{c as a,y as b,r as c,o as d,f as e,s as f,n as q,u as s};
