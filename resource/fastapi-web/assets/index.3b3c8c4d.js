import{u as x,a as D,_ as U,r as Z}from"./index.75d20446.js";/* empty css              *//* empty css               */import{d as N,e as d,r as z,t as A,B as y,aD as F,aG as o,aH as e,u as t,F as v,aL as q,aM as b,b1 as E,b2 as P,b3 as J,b4 as ee,aU as K,b5 as Q,g as W,c as j,C as G,aJ as O,b6 as ae,b7 as le,aE as T,b8 as te,b9 as oe,ba as se}from"./arco.4a860a7b.js";import{F as ne}from"./index.0b9d43f9.js";/* empty css               *//* empty css               *//* empty css               */import{f as H,g as re}from"./vue.2ee6e12c.js";import{g as X}from"./common.096b0b43.js";import{F as ce}from"./forget.198ea895.js";import"./chart.e67bf0c7.js";const ie={class:"remember-me"},ue=N({__name:"account-login",setup(I){const{proxy:m}=W(),{t:a}=x(),$=H(),w=D(),c=d(!1),l=re("login-config",{rememberMe:!0,username:"",password:""}),_=z({form:{username:l.value.username,password:l.value.password,captcha:"",uuid:""},rules:{username:[{required:!0,message:a("login.account.error.required.username")}],password:[{required:!0,message:a("login.account.error.required.password")}],captcha:[{required:!0,message:a("login.account.error.required.captcha")}]}}),{form:i,rules:g}=A(_),p=({errors:n,values:r})=>{c.value||n||(c.value=!0,w.login({account:r.username,password:r.password,terminal:"web",channel:"user",method:"account"}).then(()=>{window.localStorage.setItem("userRole","user");const{redirect:C,...V}=$.currentRoute.value.query;$.push({name:C||"Workplace",query:{...V}});const{rememberMe:R}=l.value,{username:u}=r;l.value.username=R?u:"",m.$message.success(a("login.success"))}).catch(()=>{}).finally(()=>{c.value=!1}))},L=n=>{l.value.rememberMe=n};return(n,r)=>{const C=E,V=P,R=J,u=ee,s=K,f=Q;return y(),F(f,{ref:"formRef",model:t(i),rules:t(g),layout:"vertical",size:"large",class:"login-form",onSubmit:p},{default:o(()=>[e(V,{field:"username","hide-label":""},{default:o(()=>[e(C,{modelValue:t(i).username,"onUpdate:modelValue":r[0]||(r[0]=h=>t(i).username=h),placeholder:n.$t("login.account.placeholder.username"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),e(V,{field:"password","hide-label":""},{default:o(()=>[e(R,{modelValue:t(i).password,"onUpdate:modelValue":r[1]||(r[1]=h=>t(i).password=h),placeholder:n.$t("login.account.placeholder.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),v("div",ie,[e(u,{"model-value":t(l).rememberMe,onChange:L},{default:o(()=>[q(b(n.$t("login.rememberMe")),1)]),_:1},8,["model-value","onChange"])]),e(s,{class:"btn",loading:c.value,type:"primary","html-type":"submit"},{default:o(()=>[q(b(n.$t("login.button")),1)]),_:1},8,["loading"])]),_:1},8,["model","rules"])}}});const de=U(ue,[["__scopeId","data-v-ed25edaa"]]),me={class:"login-email-title"},pe=N({__name:"email-login",setup(I){const{proxy:m}=W(),{t:a}=x(),$=H(),w=D(),c=d(!1),l=d(!1),_=d(!1),i=d(60),g=d(),p=d("login.captcha.get"),L=j(()=>a(p.value)),n=z({form:{email:"",captcha:"",terminal:"web",channel:"login"},rules:{email:[{required:!0,message:a("login.email.error.required.email")}],captcha:[{required:!0,message:a("login.email.error.required.captcha")}]}}),{form:r,rules:C}=A(n),V=()=>{window.clearInterval(g.value),i.value=60,p.value="login.captcha.get",_.value=!1},R=()=>{l.value||m.$refs.formRef.validateField("email",s=>{s||(l.value=!0,p.value="login.captcha.ing",X({email:r.value.email,channel:"login"}).then(()=>{l.value=!1,_.value=!0,p.value=`${a("login.captcha.get")}(${i.value-=1}s)`,g.value=window.setInterval(()=>{i.value-=1,p.value=`${a("login.captcha.get")}(${i.value}s)`,i.value<=0&&V()},1e3)}).catch(()=>{V(),l.value=!1}))})},u=({errors:s,values:f})=>{c.value||s||(c.value=!0,w.login({account:f.email,code:f.captcha,terminal:"web",channel:"user",method:"code"}).then(()=>{window.localStorage.setItem("userRole","user");const{redirect:h,...S}=$.currentRoute.value.query;$.push({name:h||"Workplace",query:{...S}}),m.$message.success(a("login.success"))}).catch(()=>{r.value.captcha=""}).finally(()=>{c.value=!1}))};return(s,f)=>{const h=E,S=P,B=K,M=Q;return y(),G(O,null,[v("div",me,b(s.$t("login.email.title")),1),e(M,{ref:"formRef",model:t(r),rules:t(C),layout:"vertical",size:"large",class:"login-form",onSubmit:u},{default:o(()=>[e(S,{field:"email","hide-label":""},{default:o(()=>[e(h,{modelValue:t(r).email,"onUpdate:modelValue":f[0]||(f[0]=k=>t(r).email=k),placeholder:s.$t("login.email.placeholder.email"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),e(S,{field:"captcha","hide-label":""},{default:o(()=>[e(h,{modelValue:t(r).captcha,"onUpdate:modelValue":f[1]||(f[1]=k=>t(r).captcha=k),placeholder:s.$t("login.email.placeholder.captcha"),"max-length":6,"allow-clear":"",style:{flex:"1 1"}},null,8,["modelValue","placeholder"]),e(B,{class:"captcha-btn",loading:l.value,disabled:_.value,onClick:R},{default:o(()=>[q(b(t(L)),1)]),_:1},8,["loading","disabled"])]),_:1}),e(B,{class:"btn",loading:c.value,type:"primary","html-type":"submit"},{default:o(()=>[q(b(s.$t("login.button")),1)]),_:1},8,["loading"])]),_:1},8,["model","rules"])],64)}}});const _e=U(pe,[["__scopeId","data-v-39becfa2"]]),ge={class:"sub-title"},fe=N({__name:"register",setup(I){const{proxy:m}=W(),{t:a}=x(),$=H(),w=d(!1),c=d(!1),l=d(!1),_=d(60),i=d(),g=d("login.captcha.get"),p=j(()=>a(g.value)),L=z({form:{email:"",password:"",captcha:""},rules:{email:[{required:!0,message:a("login.email.error.required.email")}],password:[{required:!0,message:a("login.email.error.required.password")}],captcha:[{required:!0,message:a("login.email.error.required.captcha")}]}}),{form:n,rules:r}=A(L),C=()=>{window.clearInterval(i.value),_.value=60,g.value="login.captcha.get",l.value=!1},V=()=>{c.value||m.$refs.formRef.validateField("email",u=>{u||(c.value=!0,g.value="login.captcha.ing",X({email:n.value.email,channel:"register"}).then(()=>{c.value=!1,l.value=!0,g.value=`${a("login.captcha.get")}(${_.value-=1}s)`,i.value=window.setInterval(()=>{_.value-=1,g.value=`${a("login.captcha.get")}(${_.value}s)`,_.value<=0&&C()},1e3)}).catch(()=>{C(),c.value=!1}))})},R=({errors:u,values:s})=>{w.value||u||(w.value=!0,Z({account:s.email,password:s.password,terminal:"web",code:s.captcha}).then(()=>{m.$message.success(a("register.success")),$.go(0)}).catch(()=>{n.value.captcha=""}).finally(()=>{w.value=!1}))};return(u,s)=>{const f=E,h=P,S=K,B=J,M=Q;return y(),G(O,null,[v("div",ge,b(u.$t("login.form.register.title")),1),e(M,{ref:"formRef",model:t(n),rules:t(r),layout:"vertical",size:"large",class:"login-form",onSubmit:R},{default:o(()=>[e(h,{field:"email","hide-label":""},{default:o(()=>[e(f,{modelValue:t(n).email,"onUpdate:modelValue":s[0]||(s[0]=k=>t(n).email=k),placeholder:u.$t("login.email.placeholder.email"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),e(h,{field:"captcha","hide-label":""},{default:o(()=>[e(f,{modelValue:t(n).captcha,"onUpdate:modelValue":s[1]||(s[1]=k=>t(n).captcha=k),placeholder:u.$t("login.email.placeholder.captcha"),"max-length":6,"allow-clear":"",style:{flex:"1 1"}},null,8,["modelValue","placeholder"]),e(S,{class:"captcha-btn",loading:c.value,disabled:l.value,onClick:V},{default:o(()=>[q(b(t(p)),1)]),_:1},8,["loading","disabled"])]),_:1}),e(h,{field:"password","hide-label":""},{default:o(()=>[e(B,{modelValue:t(n).password,"onUpdate:modelValue":s[2]||(s[2]=k=>t(n).password=k),placeholder:u.$t("login.account.placeholder.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),e(S,{class:"btn",loading:w.value,type:"primary","html-type":"submit"},{default:o(()=>[q(b(u.$t("register.button")),1)]),_:1},8,["loading"])]),_:1},8,["model","rules"])],64)}}});const he=U(fe,[["__scopeId","data-v-d7bf457d"]]),Y=I=>(te("data-v-b84a6b65"),I=I(),oe(),I),ve={class:"root"},be=Y(()=>v("div",{class:"logo"},[v("img",{alt:"logo",src:"https://www.fastapi.ai/logo.png"}),v("div",{class:"logo-text"},"\u667A\u5143 Fast API")],-1)),we={class:"container"},ye=Y(()=>v("div",{class:"left-banner"},null,-1)),$e={class:"login-card"},Ve={class:"title"},ke={class:"actions"},qe={class:"footer"},Ce=N({__name:"index",setup(I){D().logout();const m=d(!1),a=d(!1),$=()=>{m.value=!1,a.value=!1},w=()=>{m.value=!0,a.value=!1},c=()=>{a.value=!0,m.value=!1};return(l,_)=>{const i=se,g=ae,p=le;return y(),G("div",ve,[be,v("div",we,[ye,v("div",$e,[v("div",Ve,b(l.$t("login.welcome")),1),m.value?(y(),F(he,{key:0})):a.value?(y(),F(ce,{key:1})):(y(),F(g,{key:2,class:"account-tab","default-active-key":"1"},{default:o(()=>[e(i,{key:"1",title:l.$t("login.account")},{default:o(()=>[e(de)]),_:1},8,["title"]),e(i,{key:"2",title:l.$t("login.email")},{default:o(()=>[e(_e)]),_:1},8,["title"])]),_:1})),v("div",ke,[a.value?T("",!0):(y(),F(p,{key:0,onClick:c},{default:o(()=>[q(b(l.$t("login.form.forget")),1)]),_:1})),m.value?T("",!0):(y(),F(p,{key:1,onClick:w},{default:o(()=>[q(b(l.$t("login.form.register")),1)]),_:1})),m.value||a.value?(y(),F(p,{key:2,onClick:$},{default:o(()=>[q(b(l.$t("login.form.login")),1)]),_:1})):T("",!0)])])]),v("div",qe,[e(ne)])])}}});const ze=U(Ce,[["__scopeId","data-v-b84a6b65"]]);export{ze as default};
