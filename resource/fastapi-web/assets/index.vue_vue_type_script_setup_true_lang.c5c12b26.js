import{u as x}from"./index.5cd97443.js";/* empty css                *//* empty css                *//* empty css                */import{d as q,e as D,B as e,C as s,aH as l,aG as a,u as t,aD as n,aM as i,aL as w,bJ as L,bK as B,bL as C,bM as S,bN as U}from"./arco.a9260898.js";import{u as N}from"./loading.1f346a94.js";import{q as b}from"./common.df364eef.js";import{e as I}from"./admin_user.40fe53aa.js";const M={style:{margin:"10px 0 30px 10px"}},T={key:1},V={key:1},j={key:1},E={key:1},G={key:1},H={key:1},J={key:1},K={key:1},z={key:1},A={key:1},F={key:1},O={key:1},P={key:1},Q={key:1},R={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},W={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},X={key:1},Y={key:1},Z={key:1},ee={key:1},ae={key:1},te={key:1},le={name:"UserDetail"},me=q({...le,props:{id:{type:String,default:""}},setup(v){const g=v,{t:u}=x(),{loading:_,setLoading:f}=N(!0),o=D({});return(async(m={id:g.id})=>{f(!0);try{const{data:p}=await I(m);o.value=p}catch{}finally{f(!1)}})(),(m,p)=>{const r=L,d=B,c=C,y=S,$=U;return e(),s("div",M,[l($,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[l(c,{label:t(u)("common.user_id")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",T,i(o.value.user_id),1))]),_:1},8,["label"]),l(c,{label:t(u)("common.account")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",V,i(o.value.account),1))]),_:1},8,["label"]),l(c,{label:t(u)("common.name")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",j,i(o.value.name),1))]),_:1},8,["label"]),l(c,{label:t(u)("common.email")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",E,i(o.value.email),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.quota")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",G,i(o.value.quota>0?`$${t(b)(o.value.quota)}`:o.value.quota<0?`-$${t(b)(-o.value.quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.used_quota")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",H,i(o.value.used_quota>0?`$${t(b)(o.value.used_quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.quota_expires_at")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",J,i(o.value.quota_expires_at||"-"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.quota_warning")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",K,i(m.$t(`operations.open.${o.value.quota_warning||o.value.warning_threshold===0&&o.value.expire_warning_threshold===0}`)),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.warning_threshold")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",z," $"+i(o.value.warning_threshold||50),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.expire_warning_threshold")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",A,i(o.value.expire_warning_threshold||3)+"\u5929 ",1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.warning_notice")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",F,i(m.$t(`dict.notice.${o.value.warning_notice||!1}`)),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.expire_warning_notice")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",O,i(m.$t(`dict.notice.${o.value.expire_warning_notice||!1}`)),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.exhaustion_notice")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",P,i(m.$t(`dict.notice.${o.value.exhaustion_notice||!1}`)),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.expire_notice")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",Q,i(m.$t(`dict.notice.${o.value.expire_notice||!1}`)),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.models"),span:2},{default:a(()=>{var k,h;return[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",R,i(((h=(k=o.value)==null?void 0:k.model_names)==null?void 0:h.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(c,{label:t(u)("common.remark")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",W,i(o.value.remark||"-"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.login_ip"),span:2},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",X,i(o.value.login_ip||"-"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.login_time")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",Y,i(o.value.login_time||"-"),1))]),_:1},8,["label"]),l(c,{label:t(u)("user.detail.label.login_domain")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",Z,i(o.value.login_domain||"-"),1))]),_:1},8,["label"]),l(c,{label:t(u)("common.status")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",ee,[o.value.status===1?(e(),n(y,{key:0,color:"green"},{default:a(()=>[w(i(m.$t(`dict.status.${o.value.status}`)),1)]),_:1})):(e(),n(y,{key:1,color:"red"},{default:a(()=>[w(i(m.$t(`dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(c,{label:t(u)("common.created_at")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",ae,i(o.value.created_at),1))]),_:1},8,["label"]),l(c,{label:t(u)("common.updated_at")},{default:a(()=>[t(_)?(e(),n(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),s("span",te,i(o.value.updated_at),1))]),_:1},8,["label"])]),_:1})])}}});export{me as _};
