import{u as el,G as ll,p as al,y as tl,i as ul,z as ol,_ as nl}from"./index.75d20446.js";/* empty css               *//* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                */import{c as ae,S as sl}from"./sortable.esm.d405ddc2.js";/* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as dl,r as be,e as s,c as X,w as il,B as o,C as F,aH as u,aG as t,aL as d,aM as _,aJ as I,aI as S,aD as c,u as T,aE as y,F as w,D as rl,n as ml,aK as cl,aF as pl,bA as fl,bB as vl,b2 as gl,bC as bl,b1 as _l,bD as yl,bE as kl,b5 as Cl,bF as El,ab as $l,aU as hl,bi as Fl,bj as wl,bl as Bl,bm as Vl,b4 as Dl,bG as Al,aT as Il,bH as Sl,bI as zl,a_ as Ul,bJ as Ml,g as ql}from"./arco.4a860a7b.js";import{u as Ll}from"./loading.e639bbf4.js";import{q as Ol,s as Nl,a as Tl,b as Hl,c as Pl,d as Rl}from"./model.4070cf47.js";import{f as Gl}from"./agent.c30570c2.js";import{q as jl}from"./corp.cb0b2e84.js";import"./chart.e67bf0c7.js";import"./vue.2ee6e12c.js";import"./base.87fcf6e2.js";const Jl={class:"container"},Kl={class:"action-icon"},Ql={class:"action-icon"},Wl={id:"tableSetting"},Xl={style:{"margin-right":"4px",cursor:"move"}},Yl={class:"title"},Zl={name:"ModelList"},xl=dl({...Zl,setup(ea){const{loading:_e,setLoading:g}=Ll(!0),{proxy:z}=ql(),ye=be({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),te=s([]);(async()=>{g(!0);try{const{data:e}=await jl();te.value=e.items}catch{}finally{g(!1)}})();const Y=s([]);(async()=>{try{const{data:e}=await Ol();Y.value=e.items}catch{}})();const ue=s([]);(async()=>{try{const{data:e}=await Gl();ue.value=e.items}catch{}})();const ke=async e=>{g(!0);try{await Nl(e),z.$message.success("\u5220\u9664\u6210\u529F"),U()}catch{}finally{g(!1)}},oe=()=>({corp:"",name:"",model:"",type:s(),status:s(),created_at:[]}),{t:r}=el(),b=s([]),p=s(oe()),B=s([]),H=s([]),Z=s("medium"),f=s([]),n=s(!0),ne=s();let se=!1;const D={current:1,pageSize:10,showTotal:!0,showPageSize:!0,pageSizeOptions:[10,50,100,500,1e3]},P=be({...D}),Ce=X(()=>[{name:r("searchTable.size.mini"),value:"mini"},{name:r("searchTable.size.small"),value:"small"},{name:r("searchTable.size.medium"),value:"medium"},{name:r("searchTable.size.large"),value:"large"}]),Ee=X(()=>[{title:r("model.columns.corp"),dataIndex:"corp_name",slotName:"corp_name",align:"center",width:110},{title:r("model.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:r("model.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:r("model.columns.prompt_price"),dataIndex:"prompt_price",slotName:"prompt_price",align:"center"},{title:r("model.columns.completion_price"),dataIndex:"completion_price",slotName:"completion_price",align:"center"},{title:r("model.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:r("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132},{title:r("model.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]),$e=X(()=>[{label:r("model.dict.type.1"),value:1},{label:r("model.dict.type.2"),value:2},{label:r("model.dict.type.3"),value:3},{label:r("model.dict.type.4"),value:4},{label:r("model.dict.type.100"),value:100}]),he=X(()=>[{label:r("model.dict.status.1"),value:1},{label:r("model.dict.status.2"),value:2}]),R=async(e={...D})=>{g(!0);try{const{data:l}=await Tl(e);b.value=l.items,P.current=e.current,P.pageSize=e.pageSize,P.total=l.paging.total,se=l.items.length===0}catch{}finally{g(!1)}},U=()=>{R({...D,...p.value})},Fe=e=>{R({...D,...p.value,current:e})},we=e=>{D.pageSize=e,R({...D,...p.value})};R();const Be=()=>{p.value=oe()},Ve=async e=>{g(!0);try{await Hl(e),z.$message.success("\u64CD\u4F5C\u6210\u529F"),U()}catch{}finally{g(!1)}},De=(e,l)=>{Z.value=e},Ae=(e,l,i)=>{e?B.value.splice(i,0,l):B.value=H.value.filter(E=>E.dataIndex!==l.dataIndex)},de=(e,l,i,E=!1)=>{const h=E?ae(e):e;return l>-1&&i>-1&&h.splice(l,1,h.splice(i,1,h[l]).pop()),h},Ie=e=>{e&&ml(()=>{const l=document.getElementById("tableSetting");new sl(l,{onEnd(i){const{oldIndex:E,newIndex:h}=i;de(B.value,E,h),de(H.value,E,h)}})})};il(()=>Ee.value,e=>{B.value=ae(e),B.value.forEach((l,i)=>{l.checked=!0}),H.value=ae(B.value)},{deep:!0,immediate:!0});const M=s(!1),ie=s(),C=s({}),re=s(),q=s(!1),G=s({}),Se=async()=>{g(!0);try{C.value.url="",C.value.key="",C.value.is_config_model_agent=!0,M.value=!0}catch{}finally{g(!1)}},ze=async e=>{var i;if(await((i=ie.value)==null?void 0:i.validate())){M.value=!0,e(!1);return}g(!0);try{await Pl(C.value),e(),window.location.reload()}catch{}finally{g(!1)}},Ue=()=>{M.value=!1},Me=async e=>{var i;if(await((i=re.value)==null?void 0:i.validate())){q.value=!0,e(!1);return}e(),v({action:"agent",value:"all",model_agents:G.value.model_agents})},qe=()=>{q.value=!1},me=s(),L=s(!1),j=s({}),Le=async e=>{var i;if(await((i=me.value)==null?void 0:i.validate())){L.value=!0,e(!1);return}e(),v({action:"forward",value:"all",target_model:j.value.target_model})},Oe=()=>{L.value=!1},ce=s(),O=s(!1),J=s({}),Ne=async e=>{var i;if(await((i=ce.value)==null?void 0:i.validate())){O.value=!0,e(!1);return}e(),v({action:"fallback",value:"all",fallback_model:J.value.fallback_model})},Te=()=>{O.value=!1},He=e=>{f.value=e,n.value=!e.length},v=e=>{if(f.value.length===0)z.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let l=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${f.value.length}\u6761\u6570\u636E?`;switch(e.action){case"agent":e.value===!0?l=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406?`:e.value===!1?l=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406?`:e.value==="all"&&(e.model_agents?l=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406\u542F\u7528\u5E76\u4F7F\u7528\u6240\u9009\u7684\u6A21\u578B\u4EE3\u7406?`:q.value=!0);break;case"forward":e.value===!0?l=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1?`:e.value===!1?l=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1?`:e.value==="all"&&(e.target_model?l=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1\u542F\u7528\u5E76\u5168\u90E8\u8F6C\u53D1\u5230\u6240\u9009\u6A21\u578B?`:L.value=!0);break;case"fallback":e.value===!0?l=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u6A21\u578B?`:e.value===!1?l=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u6A21\u578B?`:e.value==="all"&&(e.fallback_model?l=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${f.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u6A21\u578B\u542F\u7528\u5E76\u5168\u90E8\u540E\u5907\u5230\u6240\u9009\u6A21\u578B?`:O.value=!0);break;case"status":e.value===1?l=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${f.value.length}\u6761\u6570\u636E?`:l=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${f.value.length}\u6761\u6570\u636E?`;break;case"delete":l=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${f.value.length}\u6761\u6570\u636E?`;break}if(e.action==="agent"&&e.value==="all"&&!e.model_agents||e.action==="forward"&&e.value==="all"&&!e.target_model||e.action==="fallback"&&e.value==="all"&&!e.fallback_model)return;z.$modal.warning({title:"\u8B66\u544A",titleAlign:"start",content:l,hideCancel:!1,onOk:()=>{g(!0),e.ids=f.value,Rl(e).then(i=>{g(!1),z.$message.success("\u64CD\u4F5C\u6210\u529F"),U(),ne.value.selectAll(!1)})}})}};return(e,l)=>{const i=ll,E=cl,h=pl,K=fl,A=vl,k=gl,$=bl,Q=_l,Pe=yl,x=kl,N=Cl,pe=El,Re=$l,m=hl,fe=al,ve=Fl,ee=wl,Ge=tl,je=Bl,Je=Vl,Ke=ul,Qe=ol,We=Dl,Xe=Al,ge=Il,Ye=Sl,Ze=zl,W=Ul,xe=Ml;return o(),F("div",Jl,[u(h,{class:"container-breadcrumb"},{default:t(()=>[u(E,null,{default:t(()=>[u(i)]),_:1}),u(E,null,{default:t(()=>[d(_(e.$t("menu.model")),1)]),_:1}),u(E,null,{default:t(()=>[d(_(e.$t("menu.model.list")),1)]),_:1})]),_:1}),u(xe,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:t(()=>[u(x,null,{default:t(()=>[u($,{flex:1},{default:t(()=>[u(N,{model:p.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[u(x,{gutter:16},{default:t(()=>[u($,{span:8},{default:t(()=>[u(k,{field:"corp",label:e.$t("model.form.corp")},{default:t(()=>[u(A,{modelValue:p.value.corp,"onUpdate:modelValue":l[0]||(l[0]=a=>p.value.corp=a),placeholder:e.$t("model.form.selectDefault"),"allow-search":"","allow-clear":""},{default:t(()=>[(o(!0),F(I,null,S(te.value,a=>(o(),c(K,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),u($,{span:8},{default:t(()=>[u(k,{field:"name",label:e.$t("model.form.name")},{default:t(()=>[u(Q,{modelValue:p.value.name,"onUpdate:modelValue":l[1]||(l[1]=a=>p.value.name=a),placeholder:e.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),u($,{span:8},{default:t(()=>[u(k,{field:"model",label:e.$t("model.form.model")},{default:t(()=>[u(Q,{modelValue:p.value.model,"onUpdate:modelValue":l[2]||(l[2]=a=>p.value.model=a),placeholder:e.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),u($,{span:8},{default:t(()=>[u(k,{field:"type",label:e.$t("model.form.type")},{default:t(()=>[u(A,{modelValue:p.value.type,"onUpdate:modelValue":l[3]||(l[3]=a=>p.value.type=a),options:T($e),placeholder:e.$t("model.form.selectDefault"),"allow-search":"","allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),u($,{span:8},{default:t(()=>[u(k,{field:"status",label:e.$t("model.form.status")},{default:t(()=>[u(A,{modelValue:p.value.status,"onUpdate:modelValue":l[4]||(l[4]=a=>p.value.status=a),options:T(he),placeholder:e.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),u($,{span:8},{default:t(()=>[u(k,{field:"created_at",label:e.$t("model.form.created_at")},{default:t(()=>[u(Pe,{modelValue:p.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=a=>p.value.created_at=a),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),u(pe,{style:{height:"84px"},direction:"vertical"}),u($,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[u(ve,{direction:"vertical",size:18},{default:t(()=>[u(m,{type:"primary",onClick:U},{icon:t(()=>[u(Re)]),default:t(()=>[d(" "+_(e.$t("model.form.search")),1)]),_:1}),u(m,{onClick:Be},{icon:t(()=>[u(fe)]),default:t(()=>[d(" "+_(e.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),u(pe,{style:{"margin-top":"0","margin-bottom":"16px"}}),u(x,{style:{"margin-bottom":"16px"}},{default:t(()=>[u($,{span:12},{default:t(()=>[u(ve,null,{default:t(()=>[u(m,{type:"primary",onClick:l[6]||(l[6]=a=>e.$router.push({name:"ModelCreate"}))},{default:t(()=>[d(_(e.$t("model.operation.create")),1)]),_:1}),T(se)?(o(),c(m,{key:0,type:"primary",status:"success",onClick:l[7]||(l[7]=a=>Se())},{default:t(()=>[d(" \u521D\u59CB\u5316 ")]),_:1})):y("",!0),b.value.length!==0?(o(),c(m,{key:1,type:"primary",status:"warning",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[8]||(l[8]=a=>v({action:"agent",value:"all"}))},{default:t(()=>[d(" \u5168\u90E8\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:2,type:"primary",status:"success",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[9]||(l[9]=a=>v({action:"agent",value:!0}))},{default:t(()=>[d(" \u542F\u7528\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:3,type:"primary",status:"danger",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[10]||(l[10]=a=>v({action:"agent",value:!1}))},{default:t(()=>[d(" \u5173\u95ED\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:4,type:"primary",status:"warning",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[11]||(l[11]=a=>v({action:"forward",value:"all"}))},{default:t(()=>[d(" \u5168\u90E8\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:5,type:"primary",status:"success",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[12]||(l[12]=a=>v({action:"forward",value:!0}))},{default:t(()=>[d(" \u542F\u7528\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:6,type:"primary",status:"danger",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[13]||(l[13]=a=>v({action:"forward",value:!1}))},{default:t(()=>[d(" \u5173\u95ED\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:7,type:"primary",status:"warning",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[14]||(l[14]=a=>v({action:"fallback",value:"all"}))},{default:t(()=>[d(" \u5168\u90E8\u540E\u5907 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:8,type:"primary",status:"success",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[15]||(l[15]=a=>v({action:"fallback",value:!0}))},{default:t(()=>[d(" \u542F\u7528\u540E\u5907 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:9,type:"primary",status:"danger",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[16]||(l[16]=a=>v({action:"fallback",value:!1}))},{default:t(()=>[d(" \u5173\u95ED\u540E\u5907 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:10,type:"primary",status:"success",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[17]||(l[17]=a=>v({action:"status",value:1}))},{default:t(()=>[d(" \u542F\u7528 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:11,type:"primary",status:"danger",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[18]||(l[18]=a=>v({action:"status",value:2}))},{default:t(()=>[d(" \u7981\u7528 ")]),_:1},8,["disabled","title"])):y("",!0),b.value.length!==0?(o(),c(m,{key:12,type:"primary",status:"danger",disabled:n.value,title:n.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[19]||(l[19]=a=>v({action:"delete"}))},{default:t(()=>[d(" \u5220\u9664 ")]),_:1},8,["disabled","title"])):y("",!0)]),_:1})]),_:1}),u($,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:t(()=>[u(ee,{content:e.$t("searchTable.actions.refresh")},{default:t(()=>[w("div",{class:"action-icon",onClick:U},[u(fe,{size:"18"})])]),_:1},8,["content"]),u(Je,{onSelect:De},{content:t(()=>[(o(!0),F(I,null,S(T(Ce),a=>(o(),c(je,{key:a.value,value:a.value,class:rl({active:a.value===Z.value})},{default:t(()=>[w("span",null,_(a.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[u(ee,{content:e.$t("searchTable.actions.density")},{default:t(()=>[w("div",Kl,[u(Ge,{size:"18"})])]),_:1},8,["content"])]),_:1}),u(ee,{content:e.$t("searchTable.actions.columnSetting")},{default:t(()=>[u(Xe,{trigger:"click",position:"bl",onPopupVisibleChange:Ie},{content:t(()=>[w("div",Wl,[(o(!0),F(I,null,S(H.value,(a,V)=>(o(),F("div",{key:a.dataIndex,class:"setting"},[w("div",Xl,[u(Qe)]),w("div",null,[u(We,{modelValue:a.checked,"onUpdate:modelValue":le=>a.checked=le,onChange:le=>Ae(le,a,V)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),w("div",Yl,_(a.title==="#"?"\u5E8F\u5217\u53F7":a.title),1)]))),128))])]),default:t(()=>[w("div",Ql,[u(Ke,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),u(Ze,{ref_key:"tableRef",ref:ne,"row-key":"id",loading:T(_e),pagination:P,columns:B.value,data:b.value,bordered:!1,size:Z.value,"row-selection":ye,onPageChange:Fe,onPageSizeChange:we,onSelectionChange:He},{type:t(({record:a})=>[d(_(e.$t(`model.dict.type.${a.type}`)),1)]),prompt_price:t(({record:a})=>[d(_(a.billing_method===1?`$${a.prompt_price}/k`:"-"),1)]),completion_price:t(({record:a})=>[d(_(a.billing_method===1?`$${a.completion_price}/k`:`$${a.fixed_price}/\u6B21`),1)]),status:t(({record:a})=>[u(ge,{modelValue:a.status,"onUpdate:modelValue":V=>a.status=V,"checked-value":1,"unchecked-value":2,onChange:V=>Ve({id:`${a.id}`,status:Number(`${a.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:t(({record:a})=>[u(m,{type:"text",size:"small",onClick:V=>e.$router.push({name:"ModelDetail",query:{id:`${a.id}`}})},{default:t(()=>[d(_(e.$t("model.columns.operations.view")),1)]),_:2},1032,["onClick"]),u(m,{type:"text",size:"small",onClick:V=>e.$router.push({name:"ModelUpdate",query:{id:`${a.id}`}})},{default:t(()=>[d(_(e.$t("model.columns.operations.update")),1)]),_:2},1032,["onClick"]),u(Ye,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:V=>ke({id:`${a.id}`})},{default:t(()=>[u(m,{type:"text",size:"small"},{default:t(()=>[d(_(e.$t("model.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),u(W,{visible:M.value,"onUpdate:visible":l[23]||(l[23]=a=>M.value=a),title:e.$t("model.form.title.init_model"),onCancel:Ue,onBeforeOk:ze},{default:t(()=>[u(N,{ref_key:"initForm",ref:ie,model:C.value},{default:t(()=>[u(k,{field:"url",label:e.$t("model.label.url"),rules:[{required:!0,message:e.$t("model.error.url.required")}]},{default:t(()=>[u(Q,{modelValue:C.value.url,"onUpdate:modelValue":l[20]||(l[20]=a=>C.value.url=a),placeholder:e.$t("model.placeholder.url"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),u(k,{field:"key",label:e.$t("model.label.key"),rules:[{required:!0,message:e.$t("model.error.key.required")}]},{default:t(()=>[u(Q,{modelValue:C.value.key,"onUpdate:modelValue":l[21]||(l[21]=a=>C.value.key=a),placeholder:e.$t("model.placeholder.key"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),u(k,{field:"is_config_model_agent",label:e.$t("model.label.is_config_model_agent")},{default:t(()=>[u(ge,{modelValue:C.value.is_config_model_agent,"onUpdate:modelValue":l[22]||(l[22]=a=>C.value.is_config_model_agent=a)},null,8,["modelValue"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),u(W,{visible:q.value,"onUpdate:visible":l[25]||(l[25]=a=>q.value=a),title:e.$t("model.form.title.model_agent"),onCancel:qe,onBeforeOk:Me},{default:t(()=>[u(N,{ref_key:"agentForm",ref:re,model:G.value},{default:t(()=>[u(k,{field:"model_agents",label:e.$t("model.label.model_agents"),rules:[{required:!0,message:e.$t("model.error.model_agents.required")}]},{default:t(()=>[u(A,{modelValue:G.value.model_agents,"onUpdate:modelValue":l[24]||(l[24]=a=>G.value.model_agents=a),placeholder:e.$t("model.placeholder.model_agents"),"max-tag-count":15,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(o(!0),F(I,null,S(ue.value,a=>(o(),c(K,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),u(W,{visible:L.value,"onUpdate:visible":l[27]||(l[27]=a=>L.value=a),title:e.$t("model.form.title.forward"),onCancel:Oe,onBeforeOk:Le},{default:t(()=>[u(N,{ref_key:"forwardForm",ref:me,model:j.value},{default:t(()=>[u(k,{field:"target_model",label:e.$t("model.label.target_model"),rules:[{required:!0,message:e.$t("model.error.target_model.required")}]},{default:t(()=>[u(A,{modelValue:j.value.target_model,"onUpdate:modelValue":l[26]||(l[26]=a=>j.value.target_model=a),placeholder:e.$t("model.placeholder.target_model"),"allow-search":""},{default:t(()=>[(o(!0),F(I,null,S(Y.value,a=>(o(),c(K,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),u(W,{visible:O.value,"onUpdate:visible":l[29]||(l[29]=a=>O.value=a),title:e.$t("model.form.title.fallback"),onCancel:Te,onBeforeOk:Ne},{default:t(()=>[u(N,{ref_key:"fallbackForm",ref:ce,model:J.value},{default:t(()=>[u(k,{field:"fallback_model",label:e.$t("model.label.fallback_model"),rules:[{required:!0,message:e.$t("model.error.fallback_model.required")}]},{default:t(()=>[u(A,{modelValue:J.value.fallback_model,"onUpdate:modelValue":l[28]||(l[28]=a=>J.value.fallback_model=a),placeholder:e.$t("model.placeholder.fallback_model"),"allow-search":""},{default:t(()=>[(o(!0),F(I,null,S(Y.value,a=>(o(),c(K,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"])]),_:1})])}}});const Aa=nl(xl,[["__scopeId","data-v-dad0abd9"]]);export{Aa as default};
