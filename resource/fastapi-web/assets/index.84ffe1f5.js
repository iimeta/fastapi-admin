import{k as S,_ as D}from"./index.b015539b.js";/* empty css                *//* empty css                *//* empty css              *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                */import{d as Q,e as c,B as T,C as E,aH as e,aG as a,aL as u,aM as p,F as f,u as s,bT as n,aK as G,aF as j,b1 as A,b2 as H,aS as K,bQ as M,bV as O,bW as P,bO as W,aU as z,bi as J,b5 as X,bI as Y,bR as Z,g as x}from"./arco.a9260898.js";import{u as ee}from"./loading.1f346a94.js";import{q as ae}from"./common.df364eef.js";import{f as le}from"./admin_user.60c77e3e.js";import{f as oe}from"./vue.ad52ddbe.js";import"./chart.d103b168.js";const te={class:"container"},re={class:"wrapper"},ue={style:{"margin-left":"10px"}},se={class:"submit-btn"},de={name:"UserCreate"},ne=Q({...de,setup(me){const{proxy:y}=x(),h=oe(),b=c(),o=c({name:"",account:"",password:"",terminal:"web",quota:c(),quota_expires_at:"",remark:""}),{loading:g,setLoading:v}=ee(!1),V=async()=>{var t;if(!await((t=b.value)==null?void 0:t.validate())){v(!0);try{await le(o.value).then(()=>{y.$message.success("\u65B0\u5EFA\u6210\u529F"),h.push({name:"UserList"})})}catch{}finally{v(!1)}}},w=l=>{o.value.quota=l*5e5};return(l,t)=>{const $=S,i=G,k=j,_=A,m=H,C=K,d=M,U=O,B=P,I=W,q=z,F=J,R=X,L=Y,N=Z;return T(),E("div",te,[e(k,{class:"container-breadcrumb"},{default:a(()=>[e(i,null,{default:a(()=>[e($)]),_:1}),e(i,null,{default:a(()=>[u(p(l.$t("menu.user")),1)]),_:1}),e(i,null,{default:a(()=>[u(p(l.$t("menu.user.create")),1)]),_:1})]),_:1}),e(N,{loading:s(g),style:{width:"100%"}},{default:a(()=>[e(L,{class:"general-card","body-style":{padding:"0 20px 20px 20px"},bordered:!1},{default:a(()=>[f("div",re,[e(R,{ref_key:"formRef",ref:b,model:o.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:a(()=>[e(m,{field:"name",label:l.$t("user.label.name"),rules:[{required:!0,message:l.$t("user.error.name.required")},{match:/^.{1,30}$/,message:l.$t("user.error.name.pattern")}]},{default:a(()=>[e(_,{modelValue:o.value.name,"onUpdate:modelValue":t[0]||(t[0]=r=>o.value.name=r),placeholder:l.$t("user.placeholder.name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(m,{field:"account",label:l.$t("user.label.account"),rules:[{required:!0,message:l.$t("user.error.account.required")}]},{default:a(()=>[e(_,{modelValue:o.value.account,"onUpdate:modelValue":t[1]||(t[1]=r=>o.value.account=r),placeholder:l.$t("user.placeholder.account"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(m,{field:"password",label:l.$t("user.label.password"),rules:[{required:!0,message:l.$t("user.error.password.required")},{match:/^.{6,}$/,message:l.$t("user.error.password.pattern")}]},{default:a(()=>[e(_,{modelValue:o.value.password,"onUpdate:modelValue":t[2]||(t[2]=r=>o.value.password=r),placeholder:l.$t("user.placeholder.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),e(m,{field:"quota",label:l.$t("user.label.quota"),rules:[{required:!0,message:l.$t("user.error.quota.required")}]},{default:a(()=>[e(C,{modelValue:o.value.quota,"onUpdate:modelValue":t[3]||(t[3]=r=>o.value.quota=r),placeholder:l.$t("user.placeholder.quota"),precision:0,min:-9999999999999,max:9999999999999},null,8,["modelValue","placeholder"]),f("div",ue," $"+p(o.value.quota?s(ae)(o.value.quota):"0"),1)]),_:1},8,["label","rules"]),e(m,null,{default:a(()=>[e(U,{type:"button",onChange:w},{default:a(()=>[e(d,{value:1},{default:a(()=>[u(" $1 ")]),_:1}),e(d,{value:2},{default:a(()=>[u(" $2 ")]),_:1}),e(d,{value:5},{default:a(()=>[u(" $5 ")]),_:1}),e(d,{value:10},{default:a(()=>[u(" $10 ")]),_:1}),e(d,{value:20},{default:a(()=>[u(" $20 ")]),_:1}),e(d,{value:50},{default:a(()=>[u(" $50 ")]),_:1}),e(d,{value:100},{default:a(()=>[u(" $100 ")]),_:1}),e(d,{value:200},{default:a(()=>[u(" $200 ")]),_:1}),e(d,{value:500},{default:a(()=>[u(" $500 ")]),_:1}),e(d,{value:1e3},{default:a(()=>[u(" $1000 ")]),_:1})]),_:1},8,["onChange"])]),_:1}),e(m,{field:"quota_expires_at",label:l.$t("user.label.quota_expires_at")},{default:a(()=>[e(B,{modelValue:o.value.quota_expires_at,"onUpdate:modelValue":t[4]||(t[4]=r=>o.value.quota_expires_at=r),placeholder:l.$t("user.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":r=>s(n)(r).isBefore(s(n)()),style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>s(n)().add(1,"day")},{label:"7",value:()=>s(n)().add(7,"day")},{label:"15",value:()=>s(n)().add(15,"day")},{label:"30",value:()=>s(n)().add(30,"day")},{label:"90",value:()=>s(n)().add(90,"day")},{label:"180",value:()=>s(n)().add(180,"day")},{label:"365",value:()=>s(n)().add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"]),e(m,{field:"remark",label:l.$t("user.label.remark")},{default:a(()=>[e(I,{modelValue:o.value.remark,"onUpdate:modelValue":t[5]||(t[5]=r=>o.value.remark=r),placeholder:l.$t("user.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(F,null,{default:a(()=>[f("div",se,[e(q,{type:"secondary",onClick:t[6]||(t[6]=r=>l.$router.push({name:"UserList"}))},{default:a(()=>[u(p(l.$t("button.cancel")),1)]),_:1}),e(q,{type:"primary",onClick:V},{default:a(()=>[u(p(l.$t("button.submit")),1)]),_:1})])]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const Ue=D(ne,[["__scopeId","data-v-14d71096"]]);export{Ue as default};
