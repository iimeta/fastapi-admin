import{u as w,_ as h,t as L}from"./index.313639ab.js";/* empty css               */import{u as B}from"./loading.156d38e2.js";import{d as k,c as q,B as n,C as d,aH as t,aG as o,aJ as A,aI as S,u as x,aD as g,aM as v,bQ as $,bR as j,bS as C,bk as I,e as Q,aL as y,aK as R,aF as z,bM as F}from"./arco.4eee3931.js";import{h as M}from"./vue.a3868d4c.js";import{d as N}from"./app.1ac2158a.js";/* empty css                *//* empty css                */import"./chart.01f9d9eb.js";const V={class:"item-container"},E={key:1},G=k({__name:"profile-item",props:{type:{type:String,default:""},renderData:{type:Object,required:!0},loading:{type:Boolean,default:!1}},setup(r){const u=r,{t:a}=w(),m=q(()=>{var l,s,i;const{renderData:e}=u,p=[];return p.push({title:a("app.detail.title.baseInfo"),data:[{label:a("app.detail.label.appId"),value:e.app_id},{label:a("app.detail.label.name"),value:e.name},{label:a("app.detail.label.remark"),value:(e==null?void 0:e.remark)||"-"},{label:a("app.detail.label.created_at"),value:e.created_at},{label:a("app.detail.label.updated_at"),value:e.updated_at}]}),p.push({title:a("app.detail.title.advanced"),data:[{label:a("app.detail.label.models"),value:((l=e==null?void 0:e.model_names)==null?void 0:l.join(`
`))||"-"},{label:a("app.detail.label.isLimitQuota"),value:a(`app.dict.isLimitQuota.${(e==null?void 0:e.is_limit_quota)||!1}`)},{label:a("app.detail.label.quota"),value:(e==null?void 0:e.quota)||"-"},{label:a("app.detail.label.ip_whitelist"),value:((s=e==null?void 0:e.ip_whitelist)==null?void 0:s.join(`
`))||"-"},{label:a("app.detail.label.ip_blacklist"),value:((i=e==null?void 0:e.ip_blacklist)==null?void 0:i.join(`
`))||"-"}]}),p});return(e,p)=>{const l=$,s=j,i=C,c=I;return n(),d("div",V,[t(c,{size:16,direction:"vertical",fill:""},{default:o(()=>[(n(!0),d(A,null,S(x(m),(_,b)=>(n(),g(i,{key:b,"label-style":{textAlign:"right",width:"200px",paddingRight:"10px",color:"rgb(var(--gray-8))"},"value-style":{width:"400px"},title:_.title,data:_.data},{value:o(({value:f})=>[r.loading?(n(),g(s,{key:0,animation:!0},{default:o(()=>[t(l,{widths:["200px"],rows:1})]),_:1})):(n(),d("span",E,v(f),1))]),_:2},1032,["label-style","title","data"]))),128))]),_:1})])}}});const H=h(G,[["__scopeId","data-v-c2275c87"]]),J={class:"container"},K={name:"AppDetail"},O=k({...K,setup(r){const{loading:u,setLoading:a}=B(!0),m=M(),e=Q({});return(async(l={id:m.query.id})=>{a(!0);try{const{data:s}=await N(l);e.value=s}catch{}finally{a(!1)}})(),(l,s)=>{const i=L,c=R,_=z,b=F,f=I;return n(),d("div",J,[t(_,{class:"container-breadcrumb"},{default:o(()=>[t(c,null,{default:o(()=>[t(i)]),_:1}),t(c,null,{default:o(()=>[y(v(l.$t("menu.app")),1)]),_:1}),t(c,null,{default:o(()=>[y(v(l.$t("menu.app.detail")),1)]),_:1})]),_:1}),t(f,{direction:"vertical",size:16,fill:""},{default:o(()=>[t(b,{class:"general-card"},{default:o(()=>[t(H,{loading:x(u),"render-data":e.value},null,8,["loading","render-data"])]),_:1})]),_:1})])}}});const ae=h(O,[["__scopeId","data-v-f1dea138"]]);export{ae as default};