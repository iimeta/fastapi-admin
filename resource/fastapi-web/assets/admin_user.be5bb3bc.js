import{b as t}from"./index.75d20446.js";import{q as a}from"./base.87fcf6e2.js";function n(e){return t.post("/api/v1/admin/user/create",e)}function u(e){return t.post("/api/v1/admin/user/page",e)}function o(e){return t.post("/api/v1/admin/user/delete",e)}function p(e){return t.get("/api/v1/admin/user/detail",{params:e,paramsSerializer:r=>a.stringify(r)})}function m(e){return t.post("/api/v1/admin/user/update",e)}function d(e){return t.post("/api/v1/admin/user/change/quota/expire",e)}function c(e){return t.post("/api/v1/admin/user/change/status",e)}function f(e){return t.post("/api/v1/admin/user/grant/quota",e)}function g(e){return t.post("/api/v1/admin/user/models",e)}export{d as a,c as b,f as c,g as d,n as e,p as f,m as g,u as q,o as s};
