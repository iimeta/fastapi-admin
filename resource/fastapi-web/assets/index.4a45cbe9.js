import{j as U,_ as C}from"./index.ff8851ab.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css                */import{d as B,e as p,B as I,C as F,aH as a,aG as r,aL as u,aM as n,F as L,u as N,aK as S,aF as R,b1 as D,b2 as K,aS as T,bK as j,aU as E,bi as G,b5 as H,bJ as J,bN as M}from"./arco.aed15247.js";import{u as z}from"./loading.b5911e1d.js";import{d as A}from"./admin_user.50b1d517.js";import{f as O}from"./vue.0cc5b64a.js";import"./chart.9aa6eafa.js";import"./base.87fcf6e2.js";const P={class:"container"},Q={class:"wrapper"},W={name:"UserCreate"},X=B({...W,setup(Y){const b=O(),c=p(),o=p({name:"",account:"",password:"",terminal:"web",quota:p(),remark:""}),{loading:$,setLoading:i}=z(!1),v=async()=>{var l;if(!await((l=c.value)==null?void 0:l.validate())){i(!0);try{await A(o.value).then(()=>{b.push({name:"UserList"})})}catch{}finally{i(!1)}}};return(e,l)=>{const _=U,d=S,w=R,m=D,s=K,V=T,q=j,f=E,g=G,h=H,y=J,k=M;return I(),F("div",P,[a(w,{class:"container-breadcrumb"},{default:r(()=>[a(d,null,{default:r(()=>[a(_)]),_:1}),a(d,null,{default:r(()=>[u(n(e.$t("menu.user")),1)]),_:1}),a(d,null,{default:r(()=>[u(n(e.$t("menu.user.create")),1)]),_:1})]),_:1}),a(k,{loading:N($),style:{width:"100%"}},{default:r(()=>[a(y,{class:"general-card",bordered:!1},{default:r(()=>[L("div",Q,[a(h,{ref_key:"formRef",ref:c,model:o.value,class:"form","label-col-props":{span:5},"wrapper-col-props":{span:18}},{default:r(()=>[a(s,{field:"name",label:e.$t("user.label.name"),rules:[{required:!0,message:e.$t("user.error.name.required")},{match:/^.{1,20}$/,message:e.$t("user.error.name.pattern")}]},{default:r(()=>[a(m,{modelValue:o.value.name,"onUpdate:modelValue":l[0]||(l[0]=t=>o.value.name=t),placeholder:e.$t("user.placeholder.name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"account",label:e.$t("user.label.account"),rules:[{required:!0,message:e.$t("user.error.account.required")}]},{default:r(()=>[a(m,{modelValue:o.value.account,"onUpdate:modelValue":l[1]||(l[1]=t=>o.value.account=t),placeholder:e.$t("user.placeholder.account"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"password",label:e.$t("user.label.password"),rules:[{required:!0,message:e.$t("user.error.password.required")},{match:/^.{6,}$/,message:e.$t("user.error.password.pattern")}]},{default:r(()=>[a(m,{modelValue:o.value.password,"onUpdate:modelValue":l[2]||(l[2]=t=>o.value.password=t),placeholder:e.$t("user.placeholder.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"quota",label:e.$t("user.label.quota"),rules:[{required:!0,message:e.$t("user.error.quota.required")}]},{default:r(()=>[a(V,{modelValue:o.value.quota,"onUpdate:modelValue":l[3]||(l[3]=t=>o.value.quota=t),placeholder:e.$t("user.placeholder.quota"),precision:0,min:-9999999999999,max:9999999999999},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(s,{field:"remark",label:e.$t("user.label.remark")},{default:r(()=>[a(q,{modelValue:o.value.remark,"onUpdate:modelValue":l[4]||(l[4]=t=>o.value.remark=t),placeholder:e.$t("user.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),a(s,null,{default:r(()=>[a(g,null,{default:r(()=>[a(f,{type:"secondary",onClick:l[5]||(l[5]=t=>e.$router.push({name:"UserList"}))},{default:r(()=>[u(n(e.$t("user.button.cancel")),1)]),_:1}),a(f,{type:"primary",onClick:v},{default:r(()=>[u(n(e.$t("user.button.submit")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])])]),_:1})]),_:1},8,["loading"])])}}});const me=C(X,[["__scopeId","data-v-031c7ac0"]]);export{me as default};
