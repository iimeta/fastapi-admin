import{a as P,z as E,_ as k}from"./index.74baba8a.js";/* empty css               *//* empty css               */import{d as C,e as T,B as f,aD as y,aG as e,aH as t,C as I,aE as D,u as B,aL as s,aM as c,bl as j,c4 as W,c5 as U,bR as A,bj as L,bL as q,b2 as O,b3 as M,bC as Q,bD as G,c6 as H,bK as J,aV as z,b6 as K,F as m,bW as X,b8 as F,c1 as Y,c2 as Z,c3 as x,c as ee,c7 as te,bJ as ae,b9 as ne,ba as le,bO as oe,a_ as re,bB as ie,bF as se,bb as ce,b7 as ue}from"./arco.d2aaf5b7.js";import{u as de}from"./loading.dbaba456.js";/* empty css               *//* empty css                *//* empty css              *//* empty css              *//* empty css               *//* empty css               */import{u as _e,c as pe}from"./index.641ca982.js";/* empty css               *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css              *//* empty css               */import"./chart.61872c57.js";import"./vue.ca65198a.js";const fe=["src"],me={key:1},ge=C({__name:"user-panel",setup(i){const n=P(),l={uid:"-2",name:"avatar.png",url:n.avatar},r=[{label:"userSetting.label.name",value:n.name},{label:"userSetting.label.certification",value:n.certification},{label:"userSetting.label.accountId",value:n.accountId},{label:"userSetting.label.phone",value:n.phone},{label:"userSetting.label.registrationDate",value:n.registrationDate}],d=T([l]),a=(_,u)=>{d.value=[u]},o=_=>{const u=new AbortController;return async function(){const{onProgress:b,onError:v,onSuccess:w,fileItem:h,name:V="file"}=_;b(20);const $=new FormData;$.append(V,h.file);const p=S=>{let R;S.total>0&&(R=S.loaded/S.total*100),b(parseInt(String(R),10),S)};try{const S=await _e($,{controller:u,onUploadProgress:p});w(S)}catch(S){v(S)}}(),{abort(){u.abort()}}};return(_,u)=>{const g=E,b=j,v=W,w=U,h=A,V=L,$=q;return f(),y($,{bordered:!1},{default:e(()=>[t(V,{size:54},{default:e(()=>[t(v,{"custom-request":o,"list-type":"picture-card","file-list":d.value,"show-upload-button":!0,"show-file-list":!1,onChange:a},{"upload-button":e(()=>[t(b,{size:100,class:"info-avatar"},{"trigger-icon":e(()=>[t(g)]),default:e(()=>[d.value.length?(f(),I("img",{key:0,src:d.value[0].url},null,8,fe)):D("",!0)]),_:1})]),_:1},8,["file-list"]),t(h,{data:B(r),column:2,align:"right",layout:"inline-horizontal","label-style":{width:"140px",fontWeight:"normal",color:"rgb(var(--gray-8))"},"value-style":{width:"200px",paddingLeft:"8px",textAlign:"left"}},{label:e(({label:p})=>[s(c(_.$t(p))+" :",1)]),value:e(({value:p,data:S})=>[S.label==="userSetting.label.certification"?(f(),y(w,{key:0,color:"green",size:"small"},{default:e(()=>[s(" \u5DF2\u8BA4\u8BC1 ")]),_:1})):(f(),I("span",me,c(p),1))]),_:1},8,["data","label-style"])]),_:1})]),_:1})}}});const be=k(ge,[["__scopeId","data-v-22d72ca8"]]);const Se=C({__name:"basic-information",setup(i){const n=T(),l=T({email:"",nickname:"",countryRegion:"",area:"",address:"",profile:""}),r=async()=>{var a;await((a=n.value)==null?void 0:a.validate())},d=async()=>{var a;await((a=n.value)==null?void 0:a.resetFields())};return(a,o)=>{const _=O,u=M,g=Q,b=G,v=H,w=J,h=z,V=L,$=K;return f(),y($,{ref_key:"formRef",ref:n,model:l.value,class:"form","label-col-props":{span:8},"wrapper-col-props":{span:16}},{default:e(()=>[t(u,{field:"email",label:a.$t("userSetting.basicInfo.form.label.email"),rules:[{required:!0,message:a.$t("userSetting.form.error.email.required")}]},{default:e(()=>[t(_,{modelValue:l.value.email,"onUpdate:modelValue":o[0]||(o[0]=p=>l.value.email=p),placeholder:a.$t("userSetting.basicInfo.placeholder.email")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(u,{field:"nickname",label:a.$t("userSetting.basicInfo.form.label.nickname"),rules:[{required:!0,message:a.$t("userSetting.form.error.nickname.required")}]},{default:e(()=>[t(_,{modelValue:l.value.nickname,"onUpdate:modelValue":o[1]||(o[1]=p=>l.value.nickname=p),placeholder:a.$t("userSetting.basicInfo.placeholder.nickname")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(u,{field:"countryRegion",label:a.$t("userSetting.basicInfo.form.label.countryRegion"),rules:[{required:!0,message:a.$t("userSetting.form.error.countryRegion.required")}]},{default:e(()=>[t(b,{modelValue:l.value.countryRegion,"onUpdate:modelValue":o[2]||(o[2]=p=>l.value.countryRegion=p),placeholder:a.$t("userSetting.basicInfo.placeholder.area")},{default:e(()=>[t(g,{value:"China"},{default:e(()=>[s("\u4E2D\u56FD")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(u,{field:"area",label:a.$t("userSetting.basicInfo.form.label.area"),rules:[{required:!0,message:a.$t("userSetting.form.error.area.required")}]},{default:e(()=>[t(v,{modelValue:l.value.area,"onUpdate:modelValue":o[3]||(o[3]=p=>l.value.area=p),placeholder:a.$t("userSetting.basicInfo.placeholder.area"),options:[{label:"\u5317\u4EAC",value:"beijing",children:[{label:"\u5317\u4EAC",value:"beijing",children:[{label:"\u671D\u9633",value:"chaoyang"}]}]}],"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(u,{field:"address",label:a.$t("userSetting.basicInfo.form.label.address")},{default:e(()=>[t(_,{modelValue:l.value.address,"onUpdate:modelValue":o[4]||(o[4]=p=>l.value.address=p),placeholder:a.$t("userSetting.basicInfo.placeholder.address")},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),t(u,{field:"profile",label:a.$t("userSetting.basicInfo.form.label.profile"),rules:[{maxLength:200,message:a.$t("userSetting.form.error.profile.maxLength")}],"row-class":"keep-margin"},{default:e(()=>[t(w,{modelValue:l.value.profile,"onUpdate:modelValue":o[5]||(o[5]=p=>l.value.profile=p),placeholder:a.$t("userSetting.basicInfo.placeholder.profile")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(u,null,{default:e(()=>[t(V,null,{default:e(()=>[t(h,{type:"primary",onClick:r},{default:e(()=>[s(c(a.$t("userSetting.save")),1)]),_:1}),t(h,{type:"secondary",onClick:d},{default:e(()=>[s(c(a.$t("userSetting.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const ye=k(Se,[["__scopeId","data-v-86fa7001"]]);const ve={},he={class:"content"},$e={class:"operation"},Ie={class:"content"},ke={class:"operation"},Ce={class:"content"},we={class:"operation"},Ve={class:"content"},De={class:"operation"};function Te(i,n){const l=X,r=F,d=Y,a=Z,o=x;return f(),y(o,{bordered:!1},{default:e(()=>[t(a,null,{default:e(()=>[t(d,null,{avatar:e(()=>[t(l,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.form.label.password")),1)]),_:1})]),description:e(()=>[m("div",he,[t(l,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.placeholder.password")),1)]),_:1})]),m("div",$e,[t(r,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.button.update")),1)]),_:1})])]),_:1})]),_:1}),t(a,null,{default:e(()=>[t(d,null,{avatar:e(()=>[t(l,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.form.label.securityQuestion")),1)]),_:1})]),description:e(()=>[m("div",Ie,[t(l,{class:"tip"},{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.placeholder.securityQuestion")),1)]),_:1})]),m("div",ke,[t(r,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.button.settings")),1)]),_:1})])]),_:1})]),_:1}),t(a,null,{default:e(()=>[t(d,null,{avatar:e(()=>[t(l,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.form.label.phone")),1)]),_:1})]),description:e(()=>[m("div",Ce,[t(l,null,{default:e(()=>[s(" \u5DF2\u7ED1\u5B9A\uFF1A150******50 ")]),_:1})]),m("div",we,[t(r,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.button.update")),1)]),_:1})])]),_:1})]),_:1}),t(a,null,{default:e(()=>[t(d,null,{avatar:e(()=>[t(l,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.form.label.email")),1)]),_:1})]),description:e(()=>[m("div",Ve,[t(l,{class:"tip"},{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.placeholder.email")),1)]),_:1})]),m("div",De,[t(r,null,{default:e(()=>[s(c(i.$t("userSetting.SecuritySettings.button.update")),1)]),_:1})])]),_:1})]),_:1})]),_:1})}const Be=k(ve,[["render",Te],["__scopeId","data-v-52657ad6"]]),Le={key:1},qe=C({__name:"enterprise-certification",props:{enterpriseInfo:{type:Object,required:!0}},setup(i){const n=i,l=ee(()=>{const{accountType:r,status:d,time:a,legalPerson:o,certificateType:_,authenticationNumber:u,enterpriseName:g,enterpriseCertificateType:b,organizationCode:v}=n.enterpriseInfo;return[{label:"userSetting.certification.label.accountType",value:r},{label:"userSetting.certification.label.status",value:d},{label:"userSetting.certification.label.time",value:a},{label:"userSetting.certification.label.legalPerson",value:o},{label:"userSetting.certification.label.certificateType",value:_},{label:"userSetting.certification.label.authenticationNumber",value:u},{label:"userSetting.certification.label.enterpriseName",value:g},{label:"userSetting.certification.label.enterpriseCertificateType",value:b},{label:"userSetting.certification.label.organizationCode",value:v}]});return(r,d)=>{const a=F,o=U,_=A,u=q;return f(),y(u,{class:"general-card",title:r.$t("userSetting.certification.title.enterprise"),"header-style":{padding:"0px 20px 16px 20px"},bordered:!1},{extra:e(()=>[t(a,null,{default:e(()=>[s(c(r.$t("userSetting.certification.extra.enterprise")),1)]),_:1})]),default:e(()=>[t(_,{class:"card-content",data:B(l),column:3,align:"right",layout:"inline-horizontal","label-style":{fontWeight:"normal"},"value-style":{width:"200px",paddingLeft:"8px",textAlign:"left"}},{label:e(({label:g})=>[s(c(r.$t(g))+" :",1)]),value:e(({value:g,data:b})=>[b.label==="userSetting.certification.label.status"?(f(),y(o,{key:0,color:"green",size:"small"},{default:e(()=>[s(" \u5DF2\u8BA4\u8BC1 ")]),_:1})):(f(),I("span",Le,c(g),1))]),_:1},8,["data"])]),_:1},8,["title"])}}});const Re=k(qe,[["__scopeId","data-v-d2e6bc29"]]),N=i=>(ne("data-v-90db956b"),i=i(),le(),i),Ue={key:0},Ae=N(()=>m("span",{class:"circle"},null,-1)),ze={key:1},Fe=N(()=>m("span",{class:"circle pass"},null,-1)),Ne=C({__name:"certification-records",props:{renderData:{type:Array,default(){return[]}}},setup(i){return(n,l)=>{const r=te,d=z,a=L,o=ae,_=q;return f(),y(_,{class:"general-card",title:n.$t("userSetting.certification.title.record"),"header-style":{border:"none"},bordered:!1},{default:e(()=>[i.renderData.length?(f(),y(o,{key:0,data:i.renderData},{columns:e(()=>[t(r,{title:n.$t("userSetting.certification.columns.certificationType")},{cell:e(()=>[s(c(n.$t("userSetting.certification.cell.certificationType")),1)]),_:1},8,["title"]),t(r,{title:n.$t("userSetting.certification.columns.certificationContent"),"data-index":"certificationContent"},null,8,["title"]),t(r,{title:n.$t("userSetting.certification.columns.status")},{cell:e(({record:u})=>[u.status===0?(f(),I("p",Ue,[Ae,m("span",null,c(n.$t("userSetting.certification.cell.auditing")),1)])):D("",!0),u.status===1?(f(),I("p",ze,[Fe,m("span",null,c(n.$t("userSetting.certification.cell.pass")),1)])):D("",!0)]),_:1},8,["title"]),t(r,{title:n.$t("userSetting.certification.columns.time"),"data-index":"time"},null,8,["title"]),t(r,{title:n.$t("userSetting.certification.columns.operation")},{cell:e(({record:u})=>[t(a,null,{default:e(()=>[t(d,{type:"text"},{default:e(()=>[s(c(n.$t("userSetting.certification.button.check")),1)]),_:1}),u.status===0?(f(),y(d,{key:0,type:"text"},{default:e(()=>[s(c(n.$t("userSetting.certification.button.withdraw")),1)]),_:1})):D("",!0)]),_:2},1024)]),_:1},8,["title"])]),_:1},8,["data"])):D("",!0)]),_:1},8,["title"])}}});const Pe=k(Ne,[["__scopeId","data-v-90db956b"]]),Ee=C({__name:"certification",setup(i){const{loading:n,setLoading:l}=de(!0),r=T({enterpriseInfo:{},record:[]});return(async()=>{try{const{data:a}=await pe();r.value=a}catch{}finally{l(!1)}})(),(a,o)=>{const _=oe;return f(),y(_,{loading:B(n),style:{width:"100%"}},{default:e(()=>[t(Re,{"enterprise-info":r.value.enterpriseInfo},null,8,["enterprise-info"]),t(Pe,{"render-data":r.value.record},null,8,["render-data"])]),_:1},8,["loading"])}}}),je={class:"container"},We={name:"Setting"},Oe=C({...We,setup(i){return(n,l)=>{const r=re("Breadcrumb"),d=ie,a=se,o=ce,_=ue;return f(),I("div",je,[t(r,{items:["menu.user","menu.user.setting"]},null,8,["items"]),t(a,{style:{"margin-bottom":"16px"}},{default:e(()=>[t(d,{span:24},{default:e(()=>[t(be)]),_:1})]),_:1}),t(a,{class:"wrapper"},{default:e(()=>[t(d,{span:24},{default:e(()=>[t(_,{"default-active-key":"1",type:"rounded"},{default:e(()=>[t(o,{key:"1",title:n.$t("userSetting.tab.basicInformation")},{default:e(()=>[t(ye)]),_:1},8,["title"]),t(o,{key:"2",title:n.$t("userSetting.tab.securitySettings")},{default:e(()=>[t(Be)]),_:1},8,["title"]),t(o,{key:"3",title:n.$t("userSetting.tab.certification")},{default:e(()=>[t(Ee)]),_:1},8,["title"])]),_:1})]),_:1})]),_:1})])}}});const dt=k(Oe,[["__scopeId","data-v-6710802a"]]);export{dt as default};
