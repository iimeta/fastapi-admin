import{u as hl,J as wl,p as Bl,y as Dl,i as ql,z as Vl,_ as xl}from"./index.940e37e7.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              */import{d as Al,r as Oe,e as i,c as ee,w as zl,B as n,C as p,aH as a,aG as t,aL as o,aM as s,aJ as M,aI as Q,aD as v,u as r,aE as k,F as S,D as Sl,n as Il,aK as Ul,aF as Ml,bA as Ql,bB as Ll,b2 as Ol,bC as Hl,b1 as Nl,bD as Pl,b5 as Tl,bE as Rl,ab as jl,aU as Gl,bi as Jl,bj as Kl,bl as Wl,bm as Xl,b4 as Yl,bF as Zl,aT as et,bG as lt,bH as tt,aV as at,a_ as ut,bQ as ot,b_ as nt,bI as it,g as st}from"./arco.a9260898.js";import{u as dt}from"./loading.1f346a94.js";import{p as h,q as L}from"./common.df364eef.js";import{q as rt,s as mt,c as ct,d as pt,e as _t,f as ft}from"./model.7ab43796.js";import{c as _e,S as vt}from"./sortable.esm.a0dfbf42.js";import{f as gt}from"./agent.dafd5547.js";import{q as bt}from"./corp.6047cd2f.js";import{_ as yt}from"./index.vue_vue_type_script_setup_true_lang.00275808.js";import"./chart.d103b168.js";import"./vue.ad52ddbe.js";/* empty css                *//* empty css                */const kt={class:"container"},Ct={class:"action-icon"},$t={class:"action-icon"},Et={id:"tableSetting"},Ft={style:{"margin-right":"4px",cursor:"move"}},ht={class:"title"},wt={key:0},Bt={key:1},Dt={key:2},qt={key:3},Vt={key:4},xt={key:5},At={key:0},zt={key:1},St={key:2},It={key:3},Ut={key:4},Mt={key:5},Qt={key:6},Lt={name:"ModelList"},Ot=Al({...Lt,setup(Ht){const{loading:He,setLoading:C}=dt(!0),{proxy:H}=st(),Ne=Oe({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),fe=i([]);(async()=>{C(!0);try{const{data:l}=await bt();fe.value=l.items}catch{}finally{C(!1)}})();const le=i([]);(async()=>{try{const{data:l}=await rt();le.value=l.items}catch{}})();const te=i([]);(async()=>{try{const{data:l}=await gt();te.value=l.items}catch{}})();const Pe=async l=>{C(!0);try{await mt(l),H.$message.success("\u5220\u9664\u6210\u529F"),N()}catch{}finally{C(!1)}},ve=()=>({corp:"",name:"",model:"",type:i(),status:i(),remark:""}),{t:_}=hl(),$=i([]),g=i(ve()),I=i([]),W=i([]),ae=i("medium"),b=i([]),m=i(!0),ge=i();let be=!1;const O={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},X=Oe({...O}),Te=ee(()=>[{name:_("size.mini"),value:"mini"},{name:_("size.small"),value:"small"},{name:_("size.medium"),value:"medium"},{name:_("size.large"),value:"large"}]),Re=ee(()=>[{title:_("model.columns.corp"),dataIndex:"corp_name",slotName:"corp_name",align:"center",width:110},{title:_("model.columns.model"),dataIndex:"model",slotName:"model",align:"center",ellipsis:!0,tooltip:!0},{title:_("model.columns.prompt_price"),dataIndex:"prompt_ratio",slotName:"prompt_ratio",align:"center"},{title:_("model.columns.completion_price"),dataIndex:"completion_ratio",slotName:"completion_ratio",align:"center"},{title:_("model.columns.lb_strategy"),dataIndex:"lb_strategy",slotName:"lb_strategy",align:"center"},{title:_("model.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:_("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132},{title:_("model.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]),je=ee(()=>[{label:_("dict.model_type.1"),value:1},{label:_("dict.model_type.2"),value:2},{label:_("dict.model_type.3"),value:3},{label:_("dict.model_type.4"),value:4},{label:_("dict.model_type.100"),value:100},{label:_("dict.model_type.101"),value:101},{label:_("dict.model_type.102"),value:102}]),Ge=ee(()=>[{label:_("model.dict.status.1"),value:1},{label:_("model.dict.status.2"),value:2}]),Y=async(l={...O})=>{C(!0);try{const{data:u}=await ct(l);$.value=u.items,X.current=l.current,X.pageSize=l.pageSize,X.total=u.paging.total,be=u.items.length===0}catch{}finally{C(!1)}},N=()=>{Y({...O,...g.value})},Je=l=>{Y({...O,...g.value,current:l})},Ke=l=>{O.pageSize=l,Y({...O,...g.value})};Y();const We=()=>{g.value=ve()},Xe=async l=>{C(!0);try{await pt(l),H.$message.success("\u64CD\u4F5C\u6210\u529F"),N()}catch{}finally{C(!1)}},Ye=(l,u)=>{ae.value=l},Ze=(l,u,c)=>{l?I.value.splice(c,0,u):I.value=W.value.filter(D=>D.dataIndex!==u.dataIndex)},ye=(l,u,c,D=!1)=>{const x=D?_e(l):l;return u>-1&&c>-1&&x.splice(u,1,x.splice(c,1,x[u]).pop()),x},el=l=>{l&&Il(()=>{const u=document.getElementById("tableSetting");new vt(u,{onEnd(c){const{oldIndex:D,newIndex:x}=c;ye(I.value,D,x),ye(W.value,D,x)}})})};zl(()=>Re.value,l=>{I.value=_e(l),I.value.forEach((u,c)=>{u.checked=!0}),W.value=_e(I.value)},{deep:!0,immediate:!0});const P=i(!1),ke=i(),B=i({}),Ce=i(),T=i(!1),q=i({});q.value.lb_strategy="1";const ll=async()=>{C(!0);try{B.value.url="",B.value.key="",B.value.is_config_model_agent=!0,P.value=!0}catch{}finally{C(!1)}},tl=async l=>{var c;if(await((c=ke.value)==null?void 0:c.validate())){P.value=!0,l(!1);return}C(!0);try{await _t(B.value),l(),window.location.reload()}catch{l(!1)}finally{C(!1)}},al=()=>{P.value=!1},ul=async l=>{var c;if(await((c=Ce.value)==null?void 0:c.validate())){T.value=!0,l(!1);return}l(),y({action:"agent",value:"all",lb_strategy:q.value.lb_strategy,model_agents:q.value.model_agents})},ol=()=>{T.value=!1},$e=i(),R=i(!1),Z=i({}),nl=async l=>{var c;if(await((c=$e.value)==null?void 0:c.validate())){R.value=!0,l(!1);return}l(),y({action:"forward",value:"all",target_model:Z.value.target_model})},il=()=>{R.value=!1},Ee=i(),j=i(!1),w=i({}),sl=async l=>{var c;if(await((c=Ee.value)==null?void 0:c.validate())){j.value=!0,l(!1);return}l(),y({action:"fallback",value:"all",fallback_config:{model_agent:w.value.model_agent,model:w.value.model}})},dl=()=>{j.value=!1},rl=l=>{b.value=l,m.value=!l.length},y=l=>{if(b.value.length===0)H.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let u=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${b.value.length}\u6761\u6570\u636E?`;switch(l.action){case"agent":l.value===!0?u=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406?`:l.value===!1?u=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406?`:l.value==="all"&&(l.model_agents?u=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u4EE3\u7406\u542F\u7528\u5E76\u4F7F\u7528\u6240\u9009\u7684\u6A21\u578B\u4EE3\u7406?`:T.value=!0);break;case"forward":l.value===!0?u=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1?`:l.value===!1?u=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1?`:l.value==="all"&&(l.target_model?u=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u6A21\u578B\u8F6C\u53D1\u542F\u7528\u5E76\u5168\u90E8\u8F6C\u53D1\u5230\u6240\u9009\u6A21\u578B?`:R.value=!0);break;case"fallback":l.value===!0?u=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u6A21\u578B?`:l.value===!1?u=`\u662F\u5426\u786E\u5B9A\u5173\u95ED\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u6A21\u578B?`:l.value==="all"&&(l.fallback_config?u=`\u662F\u5426\u786E\u5B9A\u5C06\u6240\u9009${b.value.length}\u6761\u6570\u636E\u7684\u540E\u5907\u914D\u7F6E\u542F\u7528\u5E76\u5168\u90E8\u540E\u5907\u5230\u6240\u9009\u914D\u7F6E?`:j.value=!0);break;case"status":l.value===1?u=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${b.value.length}\u6761\u6570\u636E?`:u=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${b.value.length}\u6761\u6570\u636E?`;break;case"delete":u=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${b.value.length}\u6761\u6570\u636E?`;break}if(l.action==="agent"&&l.value==="all"&&!l.model_agents||l.action==="forward"&&l.value==="all"&&!l.target_model||l.action==="fallback"&&l.value==="all"&&!l.fallback_config)return;H.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:u,hideCancel:!1,onOk:()=>{C(!0),l.ids=b.value,ft(l).then(c=>{C(!1),H.$message.success("\u64CD\u4F5C\u6210\u529F"),N(),ge.value.selectAll(!1)})}})}},ue=i(!1),Fe=i([]),ml=l=>{Fe.value=l,ue.value=!0},oe=i(!1),ne=i(!1),he=i(!1),we=i([]),Be=i([]),De=i([]),qe=i([]),Ve=l=>{oe.value=!0,ne.value=!1,we.value[0]=l.text_quota,Be.value=l.image_quotas,l.search_quota>0&&(ne.value=!0,De.value[0]=l),l.search_quotas&&(he.value=!0,qe.value=l.search_quotas)},ie=i(!1),xe=i([]),Ae=l=>{xe.value[0]=l,ie.value=!0},se=i(!1),ze=i([]),Se=l=>{ze.value[0]=l,se.value=!0},de=i(!1),Ie=i(),cl=l=>{de.value=!0,Ie.value=l},pl=()=>{de.value=!1};return(l,u)=>{const c=wl,D=Ul,x=Ml,G=Ql,U=Ll,E=Ol,V=Hl,J=Nl,re=Pl,K=Tl,Ue=Rl,_l=jl,d=Gl,Me=Bl,me=Jl,ce=Kl,fl=Dl,vl=Wl,gl=Xl,bl=ql,yl=Vl,kl=Yl,Cl=Zl,Qe=et,$l=lt,A=tt,El=at,z=ut,Le=ot,f=nt,Fl=it;return n(),p("div",kt,[a(x,{class:"container-breadcrumb"},{default:t(()=>[a(D,null,{default:t(()=>[a(c)]),_:1}),a(D,null,{default:t(()=>[o(s(l.$t("menu.model")),1)]),_:1}),a(D,null,{default:t(()=>[o(s(l.$t("menu.model.list")),1)]),_:1})]),_:1}),a(Fl,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:t(()=>[a(re,null,{default:t(()=>[a(V,{flex:1},{default:t(()=>[a(K,{model:g.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[a(re,{gutter:16},{default:t(()=>[a(V,{span:8},{default:t(()=>[a(E,{field:"corp",label:l.$t("model.form.corp")},{default:t(()=>[a(U,{modelValue:g.value.corp,"onUpdate:modelValue":u[0]||(u[0]=e=>g.value.corp=e),placeholder:l.$t("model.form.selectDefault"),scrollbar:!1,"allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),p(M,null,Q(fe.value,e=>(n(),v(G,{key:e.id,value:e.id,label:e.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(V,{span:8},{default:t(()=>[a(E,{field:"name",label:l.$t("model.form.name")},{default:t(()=>[a(J,{modelValue:g.value.name,"onUpdate:modelValue":u[1]||(u[1]=e=>g.value.name=e),placeholder:l.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(V,{span:8},{default:t(()=>[a(E,{field:"model",label:l.$t("model.form.model")},{default:t(()=>[a(J,{modelValue:g.value.model,"onUpdate:modelValue":u[2]||(u[2]=e=>g.value.model=e),placeholder:l.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(V,{span:8},{default:t(()=>[a(E,{field:"type",label:l.$t("model.form.type")},{default:t(()=>[a(U,{modelValue:g.value.type,"onUpdate:modelValue":u[3]||(u[3]=e=>g.value.type=e),options:r(je),placeholder:l.$t("model.form.selectDefault"),scrollbar:!1,"allow-search":"","allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(V,{span:8},{default:t(()=>[a(E,{field:"status",label:l.$t("model.form.status")},{default:t(()=>[a(U,{modelValue:g.value.status,"onUpdate:modelValue":u[4]||(u[4]=e=>g.value.status=e),options:r(Ge),placeholder:l.$t("model.form.selectDefault"),scrollbar:!1,"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(V,{span:8},{default:t(()=>[a(E,{field:"remark",label:l.$t("model.form.remark")},{default:t(()=>[a(J,{modelValue:g.value.remark,"onUpdate:modelValue":u[5]||(u[5]=e=>g.value.remark=e),placeholder:l.$t("model.form.remark.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(Ue,{style:{height:"84px"},direction:"vertical"}),a(V,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[a(me,{direction:"vertical",size:18},{default:t(()=>[a(d,{type:"primary",onClick:N},{icon:t(()=>[a(_l)]),default:t(()=>[o(" "+s(l.$t("model.form.search")),1)]),_:1}),a(d,{onClick:We},{icon:t(()=>[a(Me)]),default:t(()=>[o(" "+s(l.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(Ue,{style:{"margin-top":"0","margin-bottom":"16px"}}),a(re,{style:{"margin-bottom":"16px"}},{default:t(()=>[a(V,{span:12},{default:t(()=>[a(me,null,{default:t(()=>[a(d,{type:"primary",onClick:u[6]||(u[6]=e=>l.$router.push({name:"ModelCreate"}))},{default:t(()=>[o(s(l.$t("model.operation.create")),1)]),_:1}),r(be)?(n(),v(d,{key:0,type:"primary",status:"success",onClick:u[7]||(u[7]=e=>ll())},{default:t(()=>[o(" \u521D\u59CB\u5316 ")]),_:1})):k("",!0),$.value.length!==0?(n(),v(d,{key:1,type:"primary",status:"warning",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[8]||(u[8]=e=>y({action:"agent",value:"all"}))},{default:t(()=>[o(" \u5168\u90E8\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:2,type:"primary",status:"success",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[9]||(u[9]=e=>y({action:"agent",value:!0}))},{default:t(()=>[o(" \u542F\u7528\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:3,type:"primary",status:"danger",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[10]||(u[10]=e=>y({action:"agent",value:!1}))},{default:t(()=>[o(" \u5173\u95ED\u4EE3\u7406 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:4,type:"primary",status:"warning",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[11]||(u[11]=e=>y({action:"forward",value:"all"}))},{default:t(()=>[o(" \u5168\u90E8\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:5,type:"primary",status:"success",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[12]||(u[12]=e=>y({action:"forward",value:!0}))},{default:t(()=>[o(" \u542F\u7528\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:6,type:"primary",status:"danger",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[13]||(u[13]=e=>y({action:"forward",value:!1}))},{default:t(()=>[o(" \u5173\u95ED\u8F6C\u53D1 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:7,type:"primary",status:"warning",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[14]||(u[14]=e=>y({action:"fallback",value:"all"}))},{default:t(()=>[o(" \u5168\u90E8\u540E\u5907 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:8,type:"primary",status:"success",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[15]||(u[15]=e=>y({action:"fallback",value:!0}))},{default:t(()=>[o(" \u542F\u7528\u540E\u5907 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:9,type:"primary",status:"danger",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[16]||(u[16]=e=>y({action:"fallback",value:!1}))},{default:t(()=>[o(" \u5173\u95ED\u540E\u5907 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:10,type:"primary",status:"success",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[17]||(u[17]=e=>y({action:"status",value:1}))},{default:t(()=>[o(" \u542F\u7528 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:11,type:"primary",status:"danger",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[18]||(u[18]=e=>y({action:"status",value:2}))},{default:t(()=>[o(" \u7981\u7528 ")]),_:1},8,["disabled","title"])):k("",!0),$.value.length!==0?(n(),v(d,{key:12,type:"primary",status:"danger",disabled:m.value,title:m.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:u[19]||(u[19]=e=>y({action:"delete"}))},{default:t(()=>[o(" \u5220\u9664 ")]),_:1},8,["disabled","title"])):k("",!0)]),_:1})]),_:1}),a(V,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:t(()=>[a(ce,{content:l.$t("actions.refresh")},{default:t(()=>[S("div",{class:"action-icon",onClick:N},[a(Me,{size:"18"})])]),_:1},8,["content"]),a(gl,{onSelect:Ye},{content:t(()=>[(n(!0),p(M,null,Q(r(Te),e=>(n(),v(vl,{key:e.value,value:e.value,class:Sl({active:e.value===ae.value})},{default:t(()=>[S("span",null,s(e.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[a(ce,{content:l.$t("actions.density")},{default:t(()=>[S("div",Ct,[a(fl,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(ce,{content:l.$t("actions.column_setting")},{default:t(()=>[a(Cl,{trigger:"click",position:"bl",onPopupVisibleChange:el},{content:t(()=>[S("div",Et,[(n(!0),p(M,null,Q(W.value,(e,F)=>(n(),p("div",{key:e.dataIndex,class:"setting"},[S("div",Ft,[a(yl)]),S("div",null,[a(kl,{modelValue:e.checked,"onUpdate:modelValue":pe=>e.checked=pe,onChange:pe=>Ze(pe,e,F)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),S("div",ht,s(e.title==="#"?"\u5E8F\u5217\u53F7":e.title),1)]))),128))])]),default:t(()=>[S("div",$t,[a(bl,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(A,{ref_key:"tableRef",ref:ge,"row-key":"id",loading:r(He),pagination:X,columns:I.value,data:$.value,bordered:!1,size:ae.value,"row-selection":Ne,onPageChange:Je,onPageSizeChange:Ke,onSelectionChange:rl},{type:t(({record:e})=>[o(s(l.$t(`dict.model_type.${e.type}`)),1)]),prompt_ratio:t(({record:e})=>[e.type===5?(n(),p("span",wt,s(e.audio_quota.billing_method===1?`$${r(h)(e.audio_quota.prompt_ratio)}/k`:`$${r(L)(e.audio_quota.fixed_quota)}/\u6B21`),1)):e.type===6?(n(),p("span",Bt,"-")):e.type===100?(n(),p("span",Dt,[a(d,{type:"text",size:"small",onClick:F=>Ve(e.multimodal_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):e.type===101?(n(),p("span",qt,[a(d,{type:"text",size:"small",onClick:F=>Ae(e.realtime_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):e.type===102?(n(),p("span",Vt,[a(d,{type:"text",size:"small",onClick:F=>Se(e.multimodal_audio_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):(n(),p("span",xt,s(e.text_quota.billing_method===1?`$${r(h)(e.text_quota.prompt_ratio)}/k`:"-"),1))]),completion_ratio:t(({record:e})=>[e.type===2?(n(),p("span",At,[a(d,{type:"text",size:"small",onClick:F=>ml(e.image_quotas)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):e.type===5?(n(),p("span",zt,"-")):e.type===6?(n(),p("span",St,s(e.audio_quota.billing_method===1?`$${r(h)(e.audio_quota.completion_ratio)}/min`:`$${r(L)(e.audio_quota.fixed_quota)}/\u6B21`),1)):e.type===100?(n(),p("span",It,[a(d,{type:"text",size:"small",onClick:F=>Ve(e.multimodal_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):e.type===101?(n(),p("span",Ut,[a(d,{type:"text",size:"small",onClick:F=>Ae(e.realtime_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):e.type===102?(n(),p("span",Mt,[a(d,{type:"text",size:"small",onClick:F=>Se(e.multimodal_audio_quota)},{default:t(()=>[o(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):(n(),p("span",Qt,s(e.text_quota.billing_method===1?`$${r(h)(e.text_quota.completion_ratio)}/k`:`$${r(L)(e.text_quota.fixed_quota)}/\u6B21`),1))]),lb_strategy:t(({record:e})=>[o(s(e.is_enable_model_agent?l.$t(`dict.lb_strategy.${e.lb_strategy||1}`):"-"),1)]),status:t(({record:e})=>[a(Qe,{modelValue:e.status,"onUpdate:modelValue":F=>e.status=F,"checked-value":1,"unchecked-value":2,onChange:F=>Xe({id:`${e.id}`,status:Number(`${e.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:t(({record:e})=>[a(d,{type:"text",size:"small",onClick:F=>cl(e.id)},{default:t(()=>[o(s(l.$t("model.columns.operations.view")),1)]),_:2},1032,["onClick"]),a(d,{type:"text",size:"small",onClick:F=>l.$router.push({name:"ModelUpdate",query:{id:`${e.id}`}})},{default:t(()=>[o(s(l.$t("model.columns.operations.update")),1)]),_:2},1032,["onClick"]),a($l,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:F=>Pe({id:`${e.id}`})},{default:t(()=>[a(d,{type:"text",size:"small"},{default:t(()=>[o(s(l.$t("model.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(El,{title:l.$t("menu.model.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:de.value,onCancel:pl},{default:t(()=>[a(yt,{id:Ie.value},null,8,["id"])]),_:1},8,["title","visible"]),a(z,{visible:P.value,"onUpdate:visible":u[23]||(u[23]=e=>P.value=e),title:l.$t("model.form.title.init_model"),onCancel:al,onBeforeOk:tl},{default:t(()=>[a(K,{ref_key:"initForm",ref:ke,model:B.value},{default:t(()=>[a(E,{field:"url",label:l.$t("model.label.url"),rules:[{required:!0,message:l.$t("model.error.url.required")}]},{default:t(()=>[a(J,{modelValue:B.value.url,"onUpdate:modelValue":u[20]||(u[20]=e=>B.value.url=e),placeholder:l.$t("model.placeholder.url"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(E,{field:"key",label:l.$t("model.label.key"),rules:[{required:!0,message:l.$t("model.error.key.required")}]},{default:t(()=>[a(J,{modelValue:B.value.key,"onUpdate:modelValue":u[21]||(u[21]=e=>B.value.key=e),placeholder:l.$t("model.placeholder.key"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(E,{field:"is_config_model_agent",label:l.$t("model.label.is_config_model_agent")},{default:t(()=>[a(Qe,{modelValue:B.value.is_config_model_agent,"onUpdate:modelValue":u[22]||(u[22]=e=>B.value.is_config_model_agent=e)},null,8,["modelValue"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),a(z,{visible:T.value,"onUpdate:visible":u[27]||(u[27]=e=>T.value=e),title:l.$t("model.form.title.model_agent"),onCancel:ol,onBeforeOk:ul},{default:t(()=>[a(K,{ref_key:"agentForm",ref:Ce,model:q.value},{default:t(()=>[a(E,{field:"lb_strategy",label:l.$t("model.label.lb_strategy"),rules:[{required:!0}]},{default:t(()=>[a(me,{size:"large"},{default:t(()=>[a(Le,{modelValue:q.value.lb_strategy,"onUpdate:modelValue":u[24]||(u[24]=e=>q.value.lb_strategy=e),value:"1","default-checked":!0},{default:t(()=>[o(" \u8F6E\u8BE2 ")]),_:1},8,["modelValue"]),a(Le,{modelValue:q.value.lb_strategy,"onUpdate:modelValue":u[25]||(u[25]=e=>q.value.lb_strategy=e),value:"2"},{default:t(()=>[o("\u6743\u91CD")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label"]),a(E,{field:"model_agents",label:l.$t("model.label.model_agents"),rules:[{required:!0,message:l.$t("model.error.model_agents.required")}]},{default:t(()=>[a(U,{modelValue:q.value.model_agents,"onUpdate:modelValue":u[26]||(u[26]=e=>q.value.model_agents=e),placeholder:l.$t("model.placeholder.model_agents"),"max-tag-count":15,scrollbar:!1,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),p(M,null,Q(te.value,e=>(n(),v(G,{key:e.id,value:e.id,label:e.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),a(z,{visible:R.value,"onUpdate:visible":u[29]||(u[29]=e=>R.value=e),title:l.$t("model.form.title.forward"),onCancel:il,onBeforeOk:nl},{default:t(()=>[a(K,{ref_key:"forwardForm",ref:$e,model:Z.value},{default:t(()=>[a(E,{field:"target_model",label:l.$t("model.label.target_model"),rules:[{required:!0,message:l.$t("model.error.target_model.required")}]},{default:t(()=>[a(U,{modelValue:Z.value.target_model,"onUpdate:modelValue":u[28]||(u[28]=e=>Z.value.target_model=e),placeholder:l.$t("model.placeholder.target_model"),scrollbar:!1,"allow-search":""},{default:t(()=>[(n(!0),p(M,null,Q(le.value,e=>(n(),v(G,{key:e.id,value:e.id,label:e.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),a(z,{visible:j.value,"onUpdate:visible":u[32]||(u[32]=e=>j.value=e),title:l.$t("model.form.title.fallback"),onCancel:dl,onBeforeOk:sl},{default:t(()=>[a(K,{ref_key:"fallbackForm",ref:Ee,model:w.value},{default:t(()=>[a(E,{field:"model_agent",label:l.$t("model.label.fallback_model_agent"),rules:[{required:!w.value.model_agent&&!w.value.model||!w.value.model,message:l.$t("model.error.fallback.required")}]},{default:t(()=>[a(U,{modelValue:w.value.model_agent,"onUpdate:modelValue":u[30]||(u[30]=e=>w.value.model_agent=e),placeholder:l.$t("model.placeholder.fallback_model_agent"),scrollbar:!1,"allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),p(M,null,Q(te.value,e=>(n(),v(G,{key:e.id,value:e.id,label:e.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),a(E,{field:"model",label:l.$t("model.label.fallback_model"),rules:[{required:!w.value.model_agent&&!w.value.model||!w.value.model_agent,message:l.$t("model.error.fallback.required")}]},{default:t(()=>[a(U,{modelValue:w.value.model,"onUpdate:modelValue":u[31]||(u[31]=e=>w.value.model=e),placeholder:l.$t("model.placeholder.fallback_model"),scrollbar:!1,"allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),p(M,null,Q(le.value,e=>(n(),v(G,{key:e.id,value:e.id,label:e.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),a(z,{visible:ue.value,"onUpdate:visible":u[33]||(u[33]=e=>ue.value=e),title:l.$t("model.columns.image_price"),width:"500px","hide-cancel":"",simple:""},{default:t(()=>[a(A,{data:Fe.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u5BBD\u5EA6","data-index":"width",align:"center"}),a(f,{title:"\u9AD8\u5EA6","data-index":"height",align:"center"}),a(f,{title:"\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(L)(e.fixed_quota)}/\u5F20`),1)]),_:1}),a(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:t(({record:e})=>[o(s(e.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])]),_:1},8,["visible","title"]),a(z,{visible:oe.value,"onUpdate:visible":u[34]||(u[34]=e=>oe.value=e),title:l.$t("model.columns.multimodal_price"),width:"550px","hide-cancel":"",simple:""},{default:t(()=>[a(A,{data:we.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.prompt_ratio)}/k`),1)]),_:1}),a(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"]),a(A,{style:{"margin-top":"15px"},data:Be.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u8BC6\u56FE\u6A21\u5F0F","data-index":"mode",align:"center"}),a(f,{title:"\u8BC6\u56FE\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(L)(e.fixed_quota)}/\u5F20`),1)]),_:1}),a(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:t(({record:e})=>[o(s(e.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"]),ne.value?(n(),v(A,{key:0,style:{"margin-top":"15px"},data:De.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u641C\u7D22\u4EF7\u683C","data-index":"search_quota",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(L)(e.search_quota)}/\u6B21`),1)]),_:1})]),_:1},8,["data"])):k("",!0),he.value?(n(),v(A,{key:1,style:{"margin-top":"15px"},data:qe.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u641C\u7D22\u4E0A\u4E0B\u6587\u5927\u5C0F","data-index":"search_context_size",align:"center"}),a(f,{title:"\u641C\u7D22\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(L)(e.fixed_quota)}/\u6B21`),1)]),_:1}),a(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:t(({record:e})=>[o(s(e.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])):k("",!0)]),_:1},8,["visible","title"]),a(z,{visible:ie.value,"onUpdate:visible":u[35]||(u[35]=e=>ie.value=e),title:l.$t("model.columns.realtime_price"),width:"550px","hide-cancel":"",simple:""},{default:t(()=>[a(A,{data:xe.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.text_quota.prompt_ratio)}/k`),1)]),_:1}),a(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.text_quota.completion_ratio)}/k`),1)]),_:1}),a(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.audio_quota.prompt_ratio)}/k`),1)]),_:1}),a(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])]),_:1},8,["visible","title"]),a(z,{visible:se.value,"onUpdate:visible":u[36]||(u[36]=e=>se.value=e),title:l.$t("model.columns.multimodal_audio_price"),width:"550px","hide-cancel":"",simple:""},{default:t(()=>[a(A,{data:ze.value,pagination:!1,bordered:!1},{columns:t(()=>[a(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.text_quota.prompt_ratio)}/k`),1)]),_:1}),a(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.text_quota.completion_ratio)}/k`),1)]),_:1}),a(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.audio_quota.prompt_ratio)}/k`),1)]),_:1}),a(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:t(({record:e})=>[o(s(`$${r(h)(e.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])]),_:1},8,["visible","title"])]),_:1})])}}});const $a=xl(Ot,[["__scopeId","data-v-72f162a8"]]);export{$a as default};
