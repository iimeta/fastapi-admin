import{u as q,_ as E}from"./index.0528be47.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as H,e as s,c as T,w as A,B as u,C as w,aH as e,aG as a,aJ as G,aI as J,aD as O,aL as P,aM as R,u as j,bA as K,bB as Q,b2 as W,bC as X,b1 as Y,bD as Z,b5 as ee,bE as ae,ab as oe,aU as le,bi as te,bH as ne,bI as se}from"./arco.91d8d802.js";import{u as re}from"./loading.d8a03711.js";import{a as de}from"./model.bc455f4c.js";import{c as C}from"./sortable.esm.2649578f.js";import{q as me}from"./corp.db4573a9.js";const ce={class:"container"},pe={name:"Models"},ie=H({...pe,props:{id:{type:String,default:""},action:{type:String,default:""}},setup(x){const d=x,{loading:V,setLoading:m}=re(!0),I=()=>({corp:"",name:"",model:"",type:s(),status:s()}),_=s([]);(async()=>{m(!0);try{const{data:o}=await me();_.value=o.items}catch{}finally{m(!1)}})();const{t:p}=q(),f=s([]),n=s(I()),c=s([]),D=s([]),k=s("medium"),L=T(()=>[{title:p("model.columns.corp"),dataIndex:"corp_name",slotName:"corp_name",align:"center"},{title:p("model.columns.name"),dataIndex:"name",slotName:"name",align:"center",ellipsis:!0,tooltip:!0},{title:p("model.columns.model"),dataIndex:"model",slotName:"model",align:"center",ellipsis:!0,tooltip:!0}]),b=async(o={id:d.id,action:d.action})=>{m(!0);try{const{data:l}=await de(o);f.value=l.items}catch{}finally{m(!1)}},M=()=>{b({id:d.id,action:d.action,...n.value})};return b(),A(()=>L.value,o=>{c.value=C(o),c.value.forEach((l,v)=>{l.checked=!0}),D.value=C(c.value)},{deep:!0,immediate:!0}),(o,l)=>{const v=K,$=Q,i=W,r=X,g=Y,y=Z,B=ee,h=ae,S=oe,N=le,z=te,F=ne,U=se;return u(),w("div",ce,[e(U,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"0px"}},{default:a(()=>[e(y,null,{default:a(()=>[e(r,{flex:1},{default:a(()=>[e(B,{model:n.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(y,{gutter:16},{default:a(()=>[e(r,{span:6},{default:a(()=>[e(i,{field:"corp",label:o.$t("model.form.corp"),"label-col-props":{span:6}},{default:a(()=>[e($,{modelValue:n.value.corp,"onUpdate:modelValue":l[0]||(l[0]=t=>n.value.corp=t),placeholder:o.$t("model.form.selectDefault"),"allow-search":"","allow-clear":""},{default:a(()=>[(u(!0),w(G,null,J(_.value,t=>(u(),O(v,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(r,{span:10},{default:a(()=>[e(i,{field:"name",label:o.$t("model.form.name"),"label-col-props":{span:6}},{default:a(()=>[e(g,{modelValue:n.value.name,"onUpdate:modelValue":l[1]||(l[1]=t=>n.value.name=t),placeholder:o.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(r,{span:8},{default:a(()=>[e(i,{field:"model",label:o.$t("model.form.model")},{default:a(()=>[e(g,{modelValue:n.value.model,"onUpdate:modelValue":l[2]||(l[2]=t=>n.value.model=t),placeholder:o.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(h,{style:{height:"32px"},direction:"vertical"}),e(r,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(z,{direction:"vertical",size:18},{default:a(()=>[e(N,{type:"primary",onClick:M},{icon:a(()=>[e(S)]),default:a(()=>[P(" "+R(o.$t("model.form.search")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(h,{style:{"margin-top":"0"}}),e(F,{"row-key":"id",loading:j(V),pagination:!1,columns:c.value,data:f.value,bordered:!1,size:k.value,scroll:{y:"380px"}},null,8,["loading","columns","data","size"])]),_:1})])}}});const Ue=E(ie,[["__scopeId","data-v-2b73f986"]]);export{Ue as M};
