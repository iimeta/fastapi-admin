import{u as q}from"./index.d6462cde.js";/* empty css                *//* empty css                *//* empty css                */import{d as $,e as D,B as e,C as n,aH as l,aG as a,u as t,aD as s,aM as u,aL as h,bJ as L,bK as B,bL as C,bM as S,bN as U}from"./arco.54c7388d.js";import{u as N}from"./loading.7321a6c2.js";import{q as f}from"./common.df364eef.js";import{e as I}from"./admin_user.b1a22a84.js";const M={style:{margin:"10px 0 30px 10px"}},T={key:1},V={key:1},j={key:1},E={key:1},G={key:1},H={key:1},J={key:1},K={key:1},z={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},A={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},F={key:1},O={key:1},P={key:1},Q={key:1},R={name:"UserDetail"},se=$({...R,props:{id:{type:String,default:""}},setup(w){const g=w,{t:i}=q(),{loading:_,setLoading:b}=N(!0),o=D({});return(async(m={id:g.id})=>{b(!0);try{const{data:p}=await I(m);o.value=p}catch{}finally{b(!1)}})(),(m,p)=>{const r=L,d=B,c=C,y=S,x=U;return e(),n("div",M,[l(x,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[l(c,{label:t(i)("common.user_id")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",T,u(o.value.user_id),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.account")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",V,u(o.value.account),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.name")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",j,u(o.value.name),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.email")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",E,u(o.value.email),1))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.quota")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",G,u(o.value.quota>0?`$${t(f)(o.value.quota)}`:o.value.quota<0?`-$${t(f)(-o.value.quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.used_quota")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",H,u(o.value.used_quota>0?`$${t(f)(o.value.used_quota)}`:"$0.00"),1))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.quota_expires_at")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",J,u(o.value.quota_expires_at||"-"),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.status")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",K,[o.value.status===1?(e(),s(y,{key:0,color:"green"},{default:a(()=>[h(u(m.$t(`dict.status.${o.value.status}`)),1)]),_:1})):(e(),s(y,{key:1,color:"red"},{default:a(()=>[h(u(m.$t(`dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.models"),span:2},{default:a(()=>{var k,v;return[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",z,u(((v=(k=o.value)==null?void 0:k.model_names)==null?void 0:v.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(c,{label:t(i)("common.remark")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",A,u(o.value.remark||"-"),1))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.login_ip"),span:2},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",F,u(o.value.login_ip||"-"),1))]),_:1},8,["label"]),l(c,{label:t(i)("user.detail.label.login_time")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",O,u(o.value.login_time||"-"),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.created_at")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",P,u(o.value.created_at),1))]),_:1},8,["label"]),l(c,{label:t(i)("common.updated_at")},{default:a(()=>[t(_)?(e(),s(d,{key:0,animation:!0},{default:a(()=>[l(r,{rows:1})]),_:1})):(e(),n("span",Q,u(o.value.updated_at),1))]),_:1},8,["label"])]),_:1})])}}});export{se as _};
