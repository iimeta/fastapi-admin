import{u as x,f as R,_ as U}from"./index.dfc266c2.js";/* empty css               *//* empty css               */import{d as k,e as r,c as S,r as T,t as D,B as L,C as z,F as E,aM as g,aH as t,aG as n,u as o,aL as V,aJ as G,b1 as H,b2 as J,aU as K,b3 as M,b5 as P,g as j}from"./arco.a9260898.js";import{f as A}from"./vue.ad52ddbe.js";import{g as O}from"./common.4c01e64d.js";const Q={class:"sub-title"},W={name:"Forget"},X=k({...W,setup(Y){const{proxy:h}=j(),{t:s}=x(),y=A(),m=r(!1),c=r(!1),p=r(!1),i=r(60),_=r(),u=r("login.captcha.get"),$=S(()=>s(u.value)),F=T({form:{email:"",password:"",captcha:""},rules:{email:[{required:!0,message:s("login.email.error.required.email")}],password:[{required:!0,message:s("login.email.error.required.password")}],captcha:[{required:!0,message:s("login.email.error.required.captcha")}]}}),{form:a,rules:C}=D(F),v=()=>{window.clearInterval(_.value),i.value=60,u.value="login.captcha.get",p.value=!1},I=()=>{c.value||h.$refs.formRef.validateField("email",l=>{l||(c.value=!0,u.value="login.captcha.ing",O({email:a.value.email,channel:"forget_account",domain:window.location.hostname}).then(()=>{c.value=!1,p.value=!0,u.value=`${s("login.captcha.get")}(${i.value-=1}s)`,_.value=window.setInterval(()=>{i.value-=1,u.value=`${s("login.captcha.get")}(${i.value}s)`,i.value<=0&&v()},1e3)}).catch(()=>{v(),c.value=!1}))})},B=({errors:l,values:e})=>{m.value||l||(m.value=!0,R({account:e.email,password:e.password,terminal:"web",channel:"user",code:e.captcha,domain:window.location.hostname}).then(()=>{h.$message.success(s("forget.success")),y.go(0)}).catch(()=>{a.value.captcha=""}).finally(()=>{m.value=!1}))};return(l,e)=>{const w=H,f=J,b=K,q=M,N=P;return L(),z(G,null,[E("div",Q,g(l.$t("login.form.forget.title")),1),t(N,{ref:"formRef",model:o(a),rules:o(C),layout:"vertical",size:"large",class:"login-form",onSubmit:B},{default:n(()=>[t(f,{field:"email","hide-label":""},{default:n(()=>[t(w,{modelValue:o(a).email,"onUpdate:modelValue":e[0]||(e[0]=d=>o(a).email=d),placeholder:l.$t("login.email.placeholder.email"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),t(f,{field:"captcha","hide-label":""},{default:n(()=>[t(w,{modelValue:o(a).captcha,"onUpdate:modelValue":e[1]||(e[1]=d=>o(a).captcha=d),placeholder:l.$t("login.email.placeholder.captcha"),"max-length":6,"allow-clear":"",style:{flex:"1 1"}},null,8,["modelValue","placeholder"]),t(b,{class:"captcha-btn",loading:c.value,disabled:p.value,onClick:I},{default:n(()=>[V(g(o($)),1)]),_:1},8,["loading","disabled"])]),_:1}),t(f,{field:"password","hide-label":""},{default:n(()=>[t(q,{modelValue:o(a).password,"onUpdate:modelValue":e[2]||(e[2]=d=>o(a).password=d),placeholder:l.$t("login.account.placeholder.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1}),t(b,{class:"btn",loading:m.value,type:"primary","html-type":"submit"},{default:n(()=>[V(g(l.$t("forget.button")),1)]),_:1},8,["loading"])]),_:1},8,["model","rules"])],64)}}});const se=U(X,[["__scopeId","data-v-16ab668a"]]);export{se as F};
