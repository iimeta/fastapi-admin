import{c as R,a as L,b as ne,S as le,_ as B,u as N,k as ue}from"./index.b015539b.js";/* empty css               *//* empty css              */import{d as E,e as C,B as V,aD as q,aG as t,aH as e,C as D,aE as ie,u as $,aL as v,aM as f,bk as ce,c4 as de,bN as pe,bi as W,bI as me,r as T,c as F,t as _e,b1 as G,b2 as x,aU as H,b5 as M,g as K,F as k,aJ as J,b$ as j,b7 as Q,c5 as X,b3 as Y,a_ as Z,c6 as fe,c7 as ge,aK as be,aF as we,bC as ve,bD as he,bb as ye,b6 as Ce}from"./arco.a9260898.js";/* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css               */import{g as Se}from"./common.85d8532b.js";import"./chart.d103b168.js";import"./vue.ad52ddbe.js";function Pe(b){return R.post("/api/v1/user/update/info",b)}function $e(b){return R.post("/api/v1/user/change/password",b)}function Ie(b){return R.post("/api/v1/user/change/email",b)}function Ve(b,u){return R.post("/api/v1/user/change/avatar",b,u)}const ke=["src"],Ee=E({__name:"user-panel",setup(b){const u=L(),r=ne(),c={uid:"-2",name:"avatar.png",url:u.avatar||r.getAvatar},i=[{label:"userCenter.label.userId",value:u.user_id},{label:"userCenter.label.account",value:u.account},{label:"userCenter.label.name",value:u.name},{label:"userCenter.label.email",value:u.email},{label:"userCenter.label.createdAt",value:u.created_at}],o=C([c]),n=(p,m)=>{o.value=[m]},d=p=>{const m=new AbortController;return async function(){const{onProgress:a,onError:s,onSuccess:S,fileItem:y,name:P="file"}=p;a(20);const _=new FormData;_.append(P,y.file);const I=g=>{let l;g.total>0&&(l=g.loaded/g.total*100),a(parseInt(String(l),10),g)};try{const g=await Ve(_,{controller:m,onUploadProgress:I});S(g),u.info()}catch(g){s(g)}}(),{abort(){m.abort()}}};return(p,m)=>{const h=le,a=ce,s=de,S=pe,y=W,P=me;return V(),q(P,{bordered:!1},{default:t(()=>[e(y,{size:54},{default:t(()=>[e(s,{"custom-request":d,"list-type":"picture-card","file-list":o.value,"show-upload-button":!0,"show-file-list":!1,onChange:n},{"upload-button":t(()=>[e(a,{size:100,class:"info-avatar"},{"trigger-icon":t(()=>[e(h)]),default:t(()=>[o.value.length?(V(),D("img",{key:0,src:o.value[0].url},null,8,ke)):ie("",!0)]),_:1})]),_:1},8,["file-list"]),e(S,{data:$(i),column:2,align:"right",layout:"inline-horizontal","label-style":{width:"140px",fontWeight:"normal",color:"rgb(var(--gray-8))"},"value-style":{width:"200px",paddingLeft:"8px",textAlign:"left"}},{label:t(({label:_})=>[v(f(p.$t(_))+" :",1)]),value:t(({value:_})=>[v(f(_),1)]),_:1},8,["data","label-style"])]),_:1})]),_:1})}}});const Ue=B(Ee,[["__scopeId","data-v-aabfc986"]]),qe=E({__name:"basic-information",setup(b){const{proxy:u}=K(),{t:r}=N(),c=C(),i=L(),o=C(!1),n=T({form:{name:i.name},rules:F(()=>({name:[{required:!0,message:r("userCenter.basicInfo.form.error.required.name")}]}))}),{form:d,rules:p}=_e(n),m=()=>{o.value||c.value.validate(a=>{a||(o.value=!0,Pe({name:d.value.name}).then(()=>{i.info(),u.$message.success(r("userCenter.basicInfo.form.save.success"))}).finally(()=>{o.value=!1}))})},h=()=>{c.value.resetFields()};return(a,s)=>{const S=G,y=x,P=H,_=W,I=M;return V(),q(I,{ref_key:"formRef",ref:c,model:$(d),rules:$(p),"label-col-props":{span:8},"wrapper-col-props":{span:16},size:"large",class:"form"},{default:t(()=>[e(y,{label:a.$t("userCenter.basicInfo.form.label.name"),field:"name"},{default:t(()=>[e(S,{modelValue:$(d).name,"onUpdate:modelValue":s[0]||(s[0]=g=>$(d).name=g),placeholder:a.$t("userCenter.basicInfo.form.placeholder.name"),"max-length":30},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(y,null,{default:t(()=>[e(_,null,{default:t(()=>[e(P,{loading:o.value,type:"primary",onClick:m},{default:t(()=>[v(f(a.$t("userCenter.basicInfo.form.save")),1)]),_:1},8,["loading"]),e(P,{onClick:h},{default:t(()=>[v(f(a.$t("userCenter.basicInfo.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model","rules"])}}});const Be=B(qe,[["__scopeId","data-v-d5a4edb8"]]);const Fe={class:"tip"},Re={class:"content"},Le={class:"operation"},De=E({__name:"update-pwd",setup(b){const{proxy:u}=K(),{t:r}=N(),c=L(),i=C(),o=C(!1),n=T({oldPassword:"",newPassword:"",rePassword:""}),d=F(()=>({oldPassword:[{required:!0,message:r("userCenter.securitySettings.updatePwd.form.error.required.oldPassword")}],newPassword:[{required:!0,message:r("userCenter.securitySettings.updatePwd.form.error.required.newPassword")},{match:/^(?=.*\d)(?=.*[a-z]).{6,32}$/,message:r("userCenter.securitySettings.updatePwd.form.error.match.newPassword")},{validator:(a,s)=>{a===n.oldPassword?s(r("userCenter.securitySettings.updatePwd.form.error.validator.newPassword")):s()}}],rePassword:[{required:!0,message:r("userCenter.securitySettings.updatePwd.form.error.required.rePassword")},{validator:(a,s)=>{a!==n.newPassword?s(r("userCenter.securitySettings.updatePwd.form.error.validator.rePassword")):s()}}]})),p=()=>{o.value=!1,i.value.resetFields()},m=()=>{i.value.validate(a=>{a||$e({old_password:n.oldPassword,new_password:n.newPassword}).then(()=>{c.info(),p(),u.$message.success(r("userCenter.basicInfo.form.change.success"))})})},h=()=>{o.value=!0};return(a,s)=>{const S=j,y=Q,P=X,_=Y,I=x,g=M,l=Z;return V(),D(J,null,[e(P,null,{avatar:t(()=>[e(S,null,{default:t(()=>[v(f(a.$t("userCenter.securitySettings.password.label")),1)]),_:1})]),description:t(()=>[k("div",Fe,f(a.$t("userCenter.securitySettings.password.tip")),1),k("div",Re,[e(S,null,{default:t(()=>[v(f(a.$t("userCenter.securitySettings.content.hasBeenSet")),1)]),_:1})]),k("div",Le,[e(y,{title:a.$t("userCenter.securitySettings.button.update"),onClick:h},{default:t(()=>[v(f(a.$t("userCenter.securitySettings.button.update")),1)]),_:1},8,["title"])])]),_:1}),e(l,{title:a.$t("userCenter.securitySettings.updatePwd.modal.title"),visible:o.value,"mask-closable":!1,"esc-to-close":!1,onOk:m,onCancel:p},{default:t(()=>[e(g,{ref_key:"formRef",ref:i,model:n,rules:$(d),size:"large"},{default:t(()=>[e(I,{label:a.$t("userCenter.securitySettings.updatePwd.form.label.oldPassword"),field:"oldPassword"},{default:t(()=>[e(_,{modelValue:n.oldPassword,"onUpdate:modelValue":s[0]||(s[0]=w=>n.oldPassword=w),placeholder:a.$t("userCenter.securitySettings.updatePwd.form.placeholder.oldPassword"),"max-length":32,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(I,{label:a.$t("userCenter.securitySettings.updatePwd.form.label.newPassword"),field:"newPassword"},{default:t(()=>[e(_,{modelValue:n.newPassword,"onUpdate:modelValue":s[1]||(s[1]=w=>n.newPassword=w),placeholder:a.$t("userCenter.securitySettings.updatePwd.form.placeholder.newPassword"),"max-length":32,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(I,{label:a.$t("userCenter.securitySettings.updatePwd.form.label.rePassword"),field:"rePassword"},{default:t(()=>[e(_,{modelValue:n.rePassword,"onUpdate:modelValue":s[2]||(s[2]=w=>n.rePassword=w),placeholder:a.$t("userCenter.securitySettings.updatePwd.form.placeholder.rePassword"),"max-length":32,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model","rules"])]),_:1},8,["title","visible"])],64)}}}),ze={class:"tip"},Ae={class:"content"},Ne={class:"operation"},Te=E({__name:"update-email",setup(b){const{proxy:u}=K(),{t:r}=N(),c=L(),i=C(),o=C(60),n=C(),d=C(!1),p=C(!1),m=C(!1),h=C("userCenter.securitySettings.captcha.get"),a=F(()=>r(h.value)),s=T({email:"",code:"",password:""}),S=F(()=>({newEmail:[{required:!0,message:r("userCenter.securitySettings.updateEmail.form.error.required.newEmail")},{type:"email",message:r("userCenter.securitySettings.updateEmail.form.error.match.newEmail")}],captcha:[{required:!0,message:r("userCenter.securitySettings.form.error.required.captcha")}],currentPassword:[{required:!0,message:r("userCenter.securitySettings.updateEmail.form.error.required.currentPassword")}]})),y=()=>{window.clearInterval(n.value),o.value=60,h.value="userCenter.securitySettings.captcha.get",p.value=!1},P=()=>{d.value||i.value.validateField("newEmail",l=>{l||(d.value=!0,h.value="userCenter.securitySettings.captcha.ing",Se({email:s.email,action:"change_email",channel:c.role,domain:window.location.hostname}).then(w=>{d.value=!1,p.value=!0,h.value=`${r("userCenter.securitySettings.captcha.get")}(${o.value-=1}s)`,n.value=window.setInterval(()=>{o.value-=1,h.value=`${r("userCenter.securitySettings.captcha.get")}(${o.value}s)`,o.value<=0&&y()},1e3)}).catch(()=>{y(),d.value=!1}))})},_=()=>{m.value=!1,i.value.resetFields(),y()},I=()=>{i.value.validate(l=>{l||Ie({email:s.email,code:s.code,password:s.password}).then(()=>{_(),c.info(),u.$message.success(r("userCenter.basicInfo.form.save.success"))})})},g=()=>{m.value=!0};return(l,w)=>{const z=j,ee=Q,te=X,O=G,A=x,ae=H,se=Y,re=M,oe=Z;return V(),D(J,null,[e(te,null,{avatar:t(()=>[e(z,null,{default:t(()=>[v(f(l.$t("userCenter.securitySettings.email.label")),1)]),_:1})]),description:t(()=>[k("div",ze,f(l.$t("userCenter.securitySettings.email.tip")),1),k("div",Ae,[$(c).email?(V(),q(z,{key:0},{default:t(()=>[v(f($(c).email),1)]),_:1})):(V(),q(z,{key:1,class:"tip"},{default:t(()=>[v(f(l.$t("userCenter.securitySettings.email.content")),1)]),_:1}))]),k("div",Ne,[e(ee,{title:l.$t("userCenter.securitySettings.button.update"),onClick:g},{default:t(()=>[v(f(l.$t("userCenter.securitySettings.button.update")),1)]),_:1},8,["title"])])]),_:1}),e(oe,{title:l.$t("userCenter.securitySettings.updateEmail.modal.title"),visible:m.value,"mask-closable":!1,"esc-to-close":!1,onOk:I,onCancel:_},{default:t(()=>[e(re,{ref_key:"formRef",ref:i,model:s,rules:$(S),size:"large"},{default:t(()=>[e(A,{label:l.$t("userCenter.securitySettings.updateEmail.form.label.email"),field:"email"},{default:t(()=>[e(O,{modelValue:s.email,"onUpdate:modelValue":w[0]||(w[0]=U=>s.email=U),placeholder:l.$t("userCenter.securitySettings.updateEmail.form.placeholder.email"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(A,{label:l.$t("userCenter.securitySettings.updateEmail.form.label.captcha"),field:"code"},{default:t(()=>[e(O,{modelValue:s.code,"onUpdate:modelValue":w[1]||(w[1]=U=>s.code=U),placeholder:l.$t("userCenter.securitySettings.form.placeholder.captcha"),"max-length":6,"allow-clear":"",style:{width:"80%"}},null,8,["modelValue","placeholder"]),e(ae,{loading:d.value,type:"primary",disabled:p.value,class:"captcha-btn",onClick:P},{default:t(()=>[v(f($(a)),1)]),_:1},8,["loading","disabled"])]),_:1},8,["label"]),e(A,{label:l.$t("userCenter.securitySettings.updateEmail.form.label.currentPassword"),field:"password"},{default:t(()=>[e(se,{modelValue:s.password,"onUpdate:modelValue":w[2]||(w[2]=U=>s.password=U),placeholder:l.$t("userCenter.securitySettings.updateEmail.form.placeholder.currentPassword"),"max-length":32,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model","rules"])]),_:1},8,["title","visible"])],64)}}});const xe=B(Te,[["__scopeId","data-v-31741037"]]),Me=E({__name:"security-settings",setup(b){return(u,r)=>{const c=fe,i=ge;return V(),q(i,{bordered:!1},{default:t(()=>[e(c,null,{default:t(()=>[e(De)]),_:1}),e(c,null,{default:t(()=>[e(xe)]),_:1})]),_:1})}}});const Ke=B(Me,[["__scopeId","data-v-c0dda919"]]),Oe={class:"container"},We={name:"Setting"},Ge=E({...We,setup(b){return(u,r)=>{const c=ue,i=be,o=we,n=ve,d=he,p=ye,m=Ce;return V(),D("div",Oe,[e(o,{class:"container-breadcrumb"},{default:t(()=>[e(i,null,{default:t(()=>[e(c)]),_:1}),e(i,null,{default:t(()=>[v(f(u.$t("menu.user.center")),1)]),_:1})]),_:1}),e(d,{style:{"margin-bottom":"16px"}},{default:t(()=>[e(n,{span:24},{default:t(()=>[e(Ue)]),_:1})]),_:1}),e(d,{class:"wrapper"},{default:t(()=>[e(n,{span:24},{default:t(()=>[e(m,{"default-active-key":"1",type:"rounded"},{default:t(()=>[e(p,{key:"1",title:u.$t("userCenter.tab.basicInformation")},{default:t(()=>[e(Be)]),_:1},8,["title"]),e(p,{key:"2",title:u.$t("userCenter.tab.securitySettings")},{default:t(()=>[e(Ke)]),_:1},8,["title"])]),_:1})]),_:1})]),_:1})])}}});const gt=B(Ge,[["__scopeId","data-v-dcd276b2"]]);export{gt as default};
