import{u as $,_ as h,A as B}from"./index.313639ab.js";/* empty css               */import{u as w}from"./loading.156d38e2.js";import{d as x,c as L,B as d,C as r,aH as t,aG as l,aJ as M,aI as S,u as k,aD as y,aM as v,bQ as C,bR as R,bS as q,bk as I,e as D,aL as g,aK as F,aF as z,bM as A}from"./arco.4eee3931.js";import{h as N}from"./vue.a3868d4c.js";import{c as P}from"./model.0035c374.js";/* empty css                *//* empty css                */import"./chart.01f9d9eb.js";const V={class:"item-container"},j={key:1},E=x({__name:"profile-item",props:{type:{type:String,default:""},renderData:{type:Object,required:!0},loading:{type:Boolean,default:!1}},setup(_){const m=_,{t:a}=$(),u=L(()=>{const{renderData:e}=m,n=[];return n.push({title:a("model.detail.title.baseInfo"),data:[{label:a("model.detail.label.corp"),value:a(`model.dict.corp.${e.corp}`)},{label:a("model.detail.label.name"),value:e.name},{label:a("model.detail.label.model"),value:e.model},{label:a("model.detail.label.type"),value:a(`model.dict.type.${e.type}`)},{label:a("model.detail.label.remark"),value:(e==null?void 0:e.remark)||"-"},{label:a("model.detail.label.created_at"),value:e.created_at},{label:a("model.detail.label.updated_at"),value:e.updated_at}]}),n.push({title:a("model.detail.title.advanced"),data:[{label:a("model.detail.label.promptRatio"),value:e.prompt_ratio},{label:a("model.detail.label.completionRatio"),value:e.completion_ratio},{label:a("model.detail.label.dataFormat"),value:a(`model.dict.data_format.${e.data_format}`)},{label:a("model.detail.label.baseUrl"),value:(e==null?void 0:e.base_url)||"-"},{label:a("model.detail.label.path"),value:(e==null?void 0:e.path)||"-"},{label:a("model.detail.label.proxy"),value:(e==null?void 0:e.proxy)||"-"},{label:a("model.detail.label.isPublic"),value:a(`model.dict.is_public.${e.is_public}`)}]}),n});return(e,n)=>{const o=C,i=R,p=q,s=I;return d(),r("div",V,[t(s,{size:16,direction:"vertical",fill:""},{default:l(()=>[(d(!0),r(M,null,S(k(u),(c,b)=>(d(),y(p,{key:b,"label-style":{textAlign:"right",width:"200px",paddingRight:"10px",color:"rgb(var(--gray-8))"},"value-style":{width:"400px"},title:c.title,data:c.data},{value:l(({value:f})=>[_.loading?(d(),y(i,{key:0,animation:!0},{default:l(()=>[t(o,{widths:["200px"],rows:1})]),_:1})):(d(),r("span",j,v(f),1))]),_:2},1032,["label-style","title","data"]))),128))]),_:1})])}}});const G=h(E,[["__scopeId","data-v-1207a633"]]),H={class:"container"},J={name:"ModelDetail"},K=x({...J,setup(_){const{loading:m,setLoading:a}=w(!0),u=N(),e=D({});return(async(o={id:u.query.id})=>{a(!0);try{const{data:i}=await P(o);e.value=i}catch{}finally{a(!1)}})(),(o,i)=>{const p=B,s=F,c=z,b=A,f=I;return d(),r("div",H,[t(c,{class:"container-breadcrumb"},{default:l(()=>[t(s,null,{default:l(()=>[t(p)]),_:1}),t(s,null,{default:l(()=>[g(v(o.$t("menu.model")),1)]),_:1}),t(s,null,{default:l(()=>[g(v(o.$t("menu.model.detail")),1)]),_:1})]),_:1}),t(f,{direction:"vertical",size:16,fill:""},{default:l(()=>[t(b,{class:"general-card"},{default:l(()=>[t(G,{loading:k(m),"render-data":e.value},null,8,["loading","render-data"])]),_:1})]),_:1})])}}});const ae=h(K,[["__scopeId","data-v-83f7ae6c"]]);export{ae as default};