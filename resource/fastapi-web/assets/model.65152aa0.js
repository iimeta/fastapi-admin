import{c as t,A as i}from"./index.740287c2.js";function a(e){return t.post("/api/v1/model/init",e)}function n(e){return t.post("/api/v1/model/create",e)}function s(e){return t.post("/api/v1/model/page",e)}function u(){return t.get("/api/v1/model/list")}function l(e){return t.post("/api/v1/model/delete",e)}function p(e){return t.get("/api/v1/model/detail",{params:e,paramsSerializer:o=>i.stringify(o)})}function d(e){return t.post("/api/v1/model/update",e)}function m(e){return t.post("/api/v1/model/change/status",e)}function c(e){return t.post("/api/v1/model/batch/operate",e)}function f(){return t.get("/api/v1/model/tree")}function v(e){return t.post("/api/v1/model/permissions",e)}export{v as a,f as b,s as c,m as d,a as e,c as f,p as g,n as h,d as i,u as q,l as s};
