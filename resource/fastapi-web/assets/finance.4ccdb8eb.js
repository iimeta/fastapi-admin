import{b as r}from"./index.d6462cde.js";import{q as a}from"./base.87fcf6e2.js";function o(e){return r.get("/api/v1/finance/bill/detail",{params:e,paramsSerializer:i=>a.stringify(i)})}function l(e){return r.post("/api/v1/finance/bill/page",e)}function p(e){return r.post("/api/v1/finance/bill/export",e,{responseType:"blob"})}function s(e){return r.post("/api/v1/finance/deal/record/page",e)}export{l as a,s as b,o as q,p as s};
