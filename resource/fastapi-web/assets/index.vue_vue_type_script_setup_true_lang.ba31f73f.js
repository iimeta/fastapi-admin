import{u as x}from"./index.940e37e7.js";/* empty css                *//* empty css                *//* empty css                */import{d as D,e as $,B as e,C as i,aH as l,aG as a,u as t,aD as n,aM as d,aL as g,bJ as M,bK as L,bL as A,bM as B,bN as S}from"./arco.a9260898.js";import{u as C}from"./loading.1f346a94.js";import{c as N}from"./agent.dafd5547.js";const j={style:{margin:"10px 0 30px 10px"}},I={key:1},T={key:1},V={key:1},q={key:1},E={key:1},G={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},H={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},J={key:1},K={key:1},U={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},z={key:1},F={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},O={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},P={key:1},Q={key:1},R={key:1},W={name:"ModelAgentDetail"},se=D({...W,props:{id:{type:String,default:""}},setup(h){const v=h,{t:s}=x(),{loading:_,setLoading:f}=C(!0),o=$({});return(async(m={id:v.id})=>{f(!0);try{const{data:y}=await N(m);o.value=y}catch{}finally{f(!1)}})(),(m,y)=>{const u=M,r=L,p=A,k=B,w=S;return e(),i("div",j,[l(w,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[l(p,{label:t(s)("common.corp"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",I,d(o.value.corp_name),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.name"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",T,d(o.value.name),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.baseUrl"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",V,d(o.value.base_url),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.path"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",q,d(o.value.path||"-"),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.weight"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",E,d(o.value.weight),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.models"),span:2},{default:a(()=>{var c,b;return[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",G,d(((b=(c=o.value)==null?void 0:c.model_names)==null?void 0:b.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.fallback_models"),span:2},{default:a(()=>{var c,b;return[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",H,d(((b=(c=o.value)==null?void 0:c.fallback_model_names)==null?void 0:b.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.is_never_disable"),span:2},{default:a(()=>{var c;return[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",J,d(t(s)(`dict.${((c=o.value)==null?void 0:c.is_never_disable)||!1}`)),1))]}),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.lb_strategy"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",K,d(m.$t(`dict.lb_strategy.${o.value.lb_strategy||1}`)),1))]),_:1},8,["label"]),l(p,{label:t(s)("model.agent.detail.label.key"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",U,d(o.value.key||"-"),1))]),_:1},8,["label"]),l(p,{label:t(s)("key.detail.label.is_auto_disabled"),span:2},{default:a(()=>{var c;return[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",z,d(t(s)(`dict.${((c=o.value)==null?void 0:c.is_auto_disabled)||!1}`)),1))]}),_:1},8,["label"]),l(p,{label:t(s)("key.detail.label.auto_disabled_reason"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",F,d(o.value.auto_disabled_reason||"-"),1))]),_:1},8,["label"]),l(p,{label:t(s)("common.remark"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",O,d(o.value.remark||"-"),1))]),_:1},8,["label"]),l(p,{label:t(s)("common.status"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",P,[o.value.status===1?(e(),n(k,{key:0,color:"green"},{default:a(()=>[g(d(m.$t(`dict.status.${o.value.status}`)),1)]),_:1})):(e(),n(k,{key:1,color:"red"},{default:a(()=>[g(d(m.$t(`dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(p,{label:t(s)("common.created_at"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",Q,d(o.value.created_at),1))]),_:1},8,["label"]),l(p,{label:t(s)("common.updated_at"),span:2},{default:a(()=>[t(_)?(e(),n(r,{key:0,animation:!0},{default:a(()=>[l(u,{rows:1})]),_:1})):(e(),i("span",R,d(o.value.updated_at),1))]),_:1},8,["label"])]),_:1})])}}});export{se as _};
