import{u as j}from"./index.940e37e7.js";/* empty css                *//* empty css                *//* empty css                */import{d as B,e as K,B as t,C as _,bu as x,bv as g,F as q,aH as l,aG as a,u as e,aD as n,aM as u,aL as b,aJ as L,bJ as S,bK as C,bL as N,bM as V,bN as E}from"./arco.a9260898.js";import{u as F}from"./loading.1f346a94.js";import{q as v}from"./common.df364eef.js";import{c as I}from"./key.a2c32148.js";const J={style:{margin:"10px 0 30px 10px"}},M={key:1},T={key:1},G={key:1},H={key:1},z={key:1},A={key:1},O={key:1},P={key:1},Q={key:1},R={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},U={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},W={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},X={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Y={key:1},Z={key:1},ee={key:1},ae={style:{margin:"10px 0 30px 10px"}},te={key:1},le={key:1},oe={key:1},se={key:1},ie={key:1},ne={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},ue={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},_e={key:1},de={key:1},re={key:1},pe={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},ye={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},me={key:1},ce={key:1},ke={name:"KeyDetail"},De=B({...ke,props:{id:{type:String,default:""}},setup($){const D=$,{t:i}=j(),{loading:d,setLoading:h}=F(!0),o=K({});return(async(c={id:D.id})=>{h(!0);try{const{data:f}=await I(c);o.value=f}catch{}finally{h(!1)}})(),(c,f)=>{const r=S,p=C,y=N,k=V,w=E;return t(),_(L,null,[x(q("div",J,[l(w,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[l(y,{label:e(i)("common.key"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",M,u(o.value.key),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.app_id")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",T,u(o.value.app_id),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.user_id")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",G,u(o.value.user_id),1))]),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.is_limit_quota")},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",H,u(e(i)(`dict.${((s=o.value)==null?void 0:s.is_limit_quota)||!1}`)),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.quota")},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",z,u((s=o.value)!=null&&s.is_limit_quota?o.value.quota>0?`$${e(v)(o.value.quota)}`:"$0.00":"\u4E0D\u9650"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.used_quota")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",A,u(o.value.used_quota>0?`$${e(v)(o.value.used_quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.quota_expires_rule")},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",O,u((s=o.value)!=null&&s.is_limit_quota&&c.$t(`key.dict.quota_expires_rule.${o.value.quota_expires_rule||1}`)||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.quota_expires_at")},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",P,u((s=o.value)!=null&&s.is_limit_quota&&o.value.quota_expires_at||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.quota_expires_minutes")},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",Q,u((s=o.value)!=null&&s.is_limit_quota&&o.value.quota_expires_minutes||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.models"),span:2},{default:a(()=>{var s,m;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",R,u(((m=(s=o.value)==null?void 0:s.model_names)==null?void 0:m.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.ip_whitelist"),span:2},{default:a(()=>{var s,m;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",U,u(((m=(s=o.value)==null?void 0:s.ip_whitelist)==null?void 0:m.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.ip_blacklist")},{default:a(()=>{var s,m;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",W,u(((m=(s=o.value)==null?void 0:s.ip_blacklist)==null?void 0:m.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("common.remark")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",X,u(o.value.remark||"-"),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.status")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",Y,[o.value.status===1?(t(),n(k,{key:0,color:"green"},{default:a(()=>[b(u(c.$t(`dict.status.${o.value.status}`)),1)]),_:1})):(t(),n(k,{key:1,color:"red"},{default:a(()=>[b(u(c.$t(`dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(y,{label:e(i)("common.created_at")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",Z,u(o.value.created_at),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.updated_at")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ee,u(o.value.updated_at),1))]),_:1},8,["label"])]),_:1})],512),[[g,o.value.type===1]]),x(q("div",ae,[l(w,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[l(y,{label:e(i)("common.key"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",te,u(o.value.key),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.corp")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",le,u(o.value.corp_name),1))]),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.used_quota")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",oe,u(o.value.used_quota>0?`$${e(v)(o.value.used_quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(y,{label:e(i)("model.agent.detail.label.weight")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",se,u(o.value.weight),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.status")},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ie,[o.value.status===1?(t(),n(k,{key:0,color:"green"},{default:a(()=>[b(u(c.$t(`dict.status.${o.value.status}`)),1)]),_:1})):(t(),n(k,{key:1,color:"red"},{default:a(()=>[b(u(c.$t(`dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.bind.models"),span:2},{default:a(()=>{var s,m;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ne,u(((m=(s=o.value)==null?void 0:s.model_names)==null?void 0:m.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.model_agent_names"),span:2},{default:a(()=>{var s,m;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ue,u(((m=(s=o.value)==null?void 0:s.model_agent_names)==null?void 0:m.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.is_agents_only"),span:2},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",_e,u(e(i)(`dict.${((s=o.value)==null?void 0:s.is_agents_only)||!1}`)),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.is_never_disable"),span:2},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",de,u(e(i)(`dict.${((s=o.value)==null?void 0:s.is_never_disable)||!1}`)),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.is_auto_disabled"),span:2},{default:a(()=>{var s;return[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",re,u(e(i)(`dict.${((s=o.value)==null?void 0:s.is_auto_disabled)||!1}`)),1))]}),_:1},8,["label"]),l(y,{label:e(i)("key.detail.label.auto_disabled_reason"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",pe,u(o.value.auto_disabled_reason||"-"),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.remark"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ye,u(o.value.remark||"-"),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.created_at"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",me,u(o.value.created_at),1))]),_:1},8,["label"]),l(y,{label:e(i)("common.updated_at"),span:2},{default:a(()=>[e(d)?(t(),n(p,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(t(),_("span",ce,u(o.value.updated_at),1))]),_:1},8,["label"])]),_:1})],512),[[g,o.value.type===2]])],64)}}});export{De as _};
