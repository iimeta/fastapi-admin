import{u as w,_ as g,A as B}from"./index.74baba8a.js";/* empty css               */import{u as L}from"./loading.dbaba456.js";import{d as h,c as S,B as i,C as _,aH as a,aG as o,aJ as j,aI as q,u as x,aD as k,aM as f,bP as D,bQ as K,bR as $,bj as I,e as C,aL as v,aK as R,aF as z,bL as A}from"./arco.d2aaf5b7.js";import{h as F}from"./vue.ca65198a.js";import{b as N}from"./key.2903752a.js";/* empty css                *//* empty css                */import"./chart.61872c57.js";const P={class:"item-container"},V={key:1},E=h({__name:"profile-item",props:{type:{type:String,default:""},renderData:{type:Object,required:!0},loading:{type:Boolean,default:!1}},setup(u){const p=u,{t}=w(),m=S(()=>{var l,n,s;const{renderData:e}=p,c=[];return c.push({title:t("key.detail.title.baseInfo"),data:[{label:e.type===1?t("key.detail.label.app_id"):t("key.detail.label.corp"),value:e.type===1?e.app_id:t(`key.dict.corp.${e.corp}`)},{label:t("key.detail.label.key"),value:e.key},{label:t("key.detail.label.quota"),value:(e==null?void 0:e.quota)||"-"},{label:t("key.detail.label.remark"),value:(e==null?void 0:e.remark)||"-"},{label:t("key.detail.label.created_at"),value:e.created_at},{label:t("key.detail.label.updated_at"),value:e.updated_at}]}),c.push({title:t("key.detail.title.advanced"),data:[{label:t("key.detail.label.models"),value:((l=e==null?void 0:e.model_names)==null?void 0:l.join(`
`))||"-"},{label:e.type===1?t("key.detail.label.ip_whitelist"):"",value:e.type===1?((n=e.ip_whitelist)==null?void 0:n.join(`
`))||"-":""},{label:e.type===1?t("key.detail.label.ip_blacklist"):"",value:e.type===1?((s=e.ip_blacklist)==null?void 0:s.join(`
`))||"-":""}]}),c});return(e,c)=>{const l=D,n=K,s=$,d=I;return i(),_("div",P,[a(d,{size:16,direction:"vertical",fill:""},{default:o(()=>[(i(!0),_(j,null,q(x(m),(r,y)=>(i(),k(s,{key:y,"label-style":{textAlign:"right",width:"200px",paddingRight:"10px",color:"rgb(var(--gray-8))"},"value-style":{width:"400px"},title:r.title,data:r.data},{value:o(({value:b})=>[u.loading?(i(),k(n,{key:0,animation:!0},{default:o(()=>[a(l,{widths:["200px"],rows:1})]),_:1})):(i(),_("span",V,f(b),1))]),_:2},1032,["label-style","title","data"]))),128))]),_:1})])}}});const G=g(E,[["__scopeId","data-v-ca818025"]]),H={class:"container"},J={name:"KeyDetail"},M=h({...J,setup(u){const{loading:p,setLoading:t}=L(!0),m=F(),e=C({});return(async(l={id:m.query.id})=>{t(!0);try{const{data:n}=await N(l);e.value=n}catch{}finally{t(!1)}})(),(l,n)=>{const s=B,d=R,r=z,y=A,b=I;return i(),_("div",H,[a(r,{class:"container-breadcrumb"},{default:o(()=>[a(d,null,{default:o(()=>[a(s)]),_:1}),a(d,null,{default:o(()=>[v(f(l.$t("menu.key")),1)]),_:1}),a(d,null,{default:o(()=>[v(f(l.$t("menu.key.detail")),1)]),_:1})]),_:1}),a(b,{direction:"vertical",size:16,fill:""},{default:o(()=>[a(y,{class:"general-card",bordered:!1},{default:o(()=>[a(G,{loading:x(p),"render-data":e.value},null,8,["loading","render-data"])]),_:1})]),_:1})])}}});const te=g(M,[["__scopeId","data-v-656fddf6"]]);export{te as default};
