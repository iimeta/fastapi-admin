import{u as w,_ as g,B}from"./index.75d20446.js";/* empty css                *//* empty css               */import{d as h,c as j,B as _,C as r,aH as t,aG as n,aJ as L,aI as S,u as x,aD as k,aM as f,bO as q,bP as K,bQ as $,bi as I,e as C,aL as v,aK as A,aF as N,bJ as O}from"./arco.4a860a7b.js";import{h as z}from"./vue.2ee6e12c.js";import{u as F}from"./loading.e639bbf4.js";import{d as J}from"./key.74d2ed05.js";/* empty css                *//* empty css                */import"./chart.e67bf0c7.js";import"./base.87fcf6e2.js";const P={class:"item-container"},R={key:1},V=h({__name:"profile-item",props:{type:{type:String,default:""},renderData:{type:Object,required:!0},loading:{type:Boolean,default:!1}},setup(u){const p=u,{t:a}=w(),m=j(()=>{var l,o,d,i,s;const{renderData:e}=p,c=[];return c.push({title:a("key.detail.title.baseInfo"),data:[{label:e.type===1?a("key.detail.label.app_id"):a("key.detail.label.corp"),value:e.type===1?e.app_id:e.corp_name},{label:a("key.detail.label.key"),value:e.key},{label:a("key.detail.label.quota"),value:(e==null?void 0:e.quota)||"-"},{label:a("key.detail.label.remark"),value:(e==null?void 0:e.remark)||"-"},{label:a("key.detail.label.created_at"),value:e.created_at},{label:a("key.detail.label.updated_at"),value:e.updated_at}]}),e.type===1&&c.push({title:a("key.detail.title.advanced"),data:[{label:a("key.detail.label.models"),value:((l=e==null?void 0:e.model_names)==null?void 0:l.join(`
`))||"-"},{label:a("key.detail.label.ip_whitelist"),value:((o=e.ip_whitelist)==null?void 0:o.join(`
`))||"-"},{label:e.type===1?a("key.detail.label.ip_blacklist"):"",value:e.type===1?((d=e.ip_blacklist)==null?void 0:d.join(`
`))||"-":""}]}),e.type===2&&c.push({title:a("key.detail.title.advanced"),data:[{label:a("key.detail.label.models"),value:((i=e==null?void 0:e.model_names)==null?void 0:i.join(`
`))||"-"},{label:a("key.detail.label.modelAgentNames"),value:((s=e==null?void 0:e.model_agent_names)==null?void 0:s.join(`
`))||"-"},{label:a("key.detail.label.isAgentsOnly"),value:a(`key.dict.is_agents_only.${e.is_agents_only}`)}]}),c});return(e,c)=>{const l=q,o=K,d=$,i=I;return _(),r("div",P,[t(i,{size:16,direction:"vertical",fill:""},{default:n(()=>[(_(!0),r(L,null,S(x(m),(s,y)=>(_(),k(d,{key:y,"label-style":{textAlign:"right",width:"200px",paddingRight:"10px",color:"rgb(var(--gray-8))"},"value-style":{width:"400px"},title:s.title,data:s.data},{value:n(({value:b})=>[u.loading?(_(),k(o,{key:0,animation:!0},{default:n(()=>[t(l,{widths:["200px"],rows:1})]),_:1})):(_(),r("span",R,f(b),1))]),_:2},1032,["label-style","title","data"]))),128))]),_:1})])}}});const D=g(V,[["__scopeId","data-v-51412a3c"]]),E={class:"container"},G={name:"KeyDetail"},H=h({...G,setup(u){const{loading:p,setLoading:a}=F(!0),m=z(),e=C({});return(async(l={id:m.query.id})=>{a(!0);try{const{data:o}=await J(l);e.value=o}catch{}finally{a(!1)}})(),(l,o)=>{const d=B,i=A,s=N,y=O,b=I;return _(),r("div",E,[t(s,{class:"container-breadcrumb"},{default:n(()=>[t(i,null,{default:n(()=>[t(d)]),_:1}),t(i,null,{default:n(()=>[v(f(l.$t("menu.key")),1)]),_:1}),t(i,null,{default:n(()=>[v(f(l.$t("menu.key.detail")),1)]),_:1})]),_:1}),t(b,{direction:"vertical",size:16,fill:""},{default:n(()=>[t(y,{class:"general-card",bordered:!1},{default:n(()=>[t(D,{loading:x(p),"render-data":e.value},null,8,["loading","render-data"])]),_:1})]),_:1})])}}});const le=g(H,[["__scopeId","data-v-96619ab4"]]);export{le as default};
