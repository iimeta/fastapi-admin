import{u as Re,C as je,p as Ge,y as Je,i as Qe,z as We,_ as Xe}from"./index.b015539b.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css              */import{d as Ye,r as ie,e as n,c as j,w as de,B as c,C as k,aH as a,aG as l,aL as s,aM as d,aJ as I,aI as A,aD as N,u as M,F as b,D as Ze,g as et,n as tt,aK as at,aF as lt,bA as ot,bB as ut,b2 as nt,bC as st,b1 as it,bD as dt,b5 as rt,bE as ct,ab as mt,aU as pt,bi as _t,bj as vt,bl as ft,bm as yt,b4 as gt,bF as kt,ad as bt,aT as ht,bG as Ct,bH as Et,aV as $t,a_ as wt,bI as Ft}from"./arco.a9260898.js";import{h as Bt,u as Dt}from"./vue.ad52ddbe.js";import{u as Vt}from"./loading.1f346a94.js";import{q as It}from"./common.df364eef.js";import{s as At,q as xt,a as St,b as zt,c as qt}from"./key.69e0d4e7.js";import{c as G,S as Lt}from"./sortable.esm.a0dfbf42.js";import{q as Nt}from"./model.c2d70173.js";import{f as Mt}from"./agent.c0347aea.js";import{q as Ut}from"./corp.7c65ca36.js";import{M as Kt}from"./models.465bd85b.js";import{_ as Pt}from"./index.vue_vue_type_script_setup_true_lang.d92b661e.js";import"./chart.d103b168.js";/* empty css                *//* empty css                */const Ot={class:"container"},Ht={class:"action-icon"},Tt={class:"action-icon"},Rt={id:"tableSetting"},jt={style:{"margin-right":"4px",cursor:"move"}},Gt={class:"title"},Jt={key:0},Qt={key:1},Wt={name:"KeyList"},Xt=Ye({...Wt,setup(Yt){const{loading:re,setLoading:p}=Vt(!0),{proxy:w}=et(),x=Bt(),ce=ie({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),J=n([]);(async()=>{p(!0);try{const{data:e}=await Ut();J.value=e.items}catch{}finally{p(!1)}})();const Q=n([]);(async()=>{try{const{data:e}=await Nt();Q.value=e.items}catch{}})();const W=n([]);(async()=>{try{const{data:e}=await Mt();W.value=e.items;const o=new Array(0);x.query.agent_id&&(o[0]=x.query.agent_id,u.value.model_agents=o)}catch{}})();const me=async e=>{p(!0);try{await At(e),w.$message.success("\u5220\u9664\u6210\u529F"),V()}catch{}finally{p(!1)}},X=()=>({type:2,corp:"",key:"",models:[],model_agents:n(),quota:n(),status:n(),remark:""}),{t:i}=Re(),Y=n([]),u=n(X()),h=n([]),S=n([]),U=n("medium"),C=n([]),E=n(!0),y=n(!0),Z=n(),F={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},$=ie({...F}),pe=j(()=>[{name:i("size.mini"),value:"mini"},{name:i("size.small"),value:"small"},{name:i("size.medium"),value:"medium"},{name:i("size.large"),value:"large"}]),_e=j(()=>[{title:i("key.columns.corp"),dataIndex:"corp_name",slotName:"corp_name",align:"center",width:110},{title:i("key.columns.key"),dataIndex:"key",slotName:"key",align:"center",width:230},{title:i("key.columns.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:i("key.columns.model_agents"),dataIndex:"model_agent_names",slotName:"model_agent_names",align:"center",ellipsis:!0,tooltip:!0},{title:i("key.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:i("common.weight"),dataIndex:"weight",slotName:"weight",align:"center",width:60},{title:i("key.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:i("key.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:i("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132},{title:i("key.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]),ve=j(()=>[{label:i("key.dict.status.1"),value:1},{label:i("key.dict.status.2"),value:2}]),ee=new Array(0);x.query.agent_id&&(ee[0]=x.query.agent_id);const z=async(e={...F,type:2,model_agents:ee})=>{p(!0);try{const{data:o}=await xt(e);Y.value=o.items,$.current=e.current,$.pageSize=e.pageSize,$.total=o.paging.total,o.items.length>0&&(u.value.corp||u.value.key||u.value.models.length>0||u.value.model_agents||u.value.status||u.value.remark)?y.value=!1:y.value=!0}catch{}finally{p(!1)}},V=()=>{z({...F,...u.value})},fe=e=>{z({...F,...u.value,current:e})},ye=e=>{F.pageSize=e,z({...F,...u.value})};z();const ge=()=>{u.value=X()},ke=async e=>{p(!0);try{await St(e),w.$message.success("\u64CD\u4F5C\u6210\u529F"),V()}catch{}finally{p(!1)}},be=(e,o)=>{U.value=e},he=(e,o,_)=>{e?h.value.splice(_,0,o):h.value=S.value.filter(v=>v.dataIndex!==o.dataIndex)},te=(e,o,_,v=!1)=>{const g=v?G(e):e;return o>-1&&_>-1&&g.splice(o,1,g.splice(_,1,g[o]).pop()),g},Ce=e=>{e&&tt(()=>{const o=document.getElementById("tableSetting");new Lt(o,{onEnd(_){const{oldIndex:v,newIndex:g}=_;te(h.value,v,g),te(S.value,v,g)}})})};de(()=>_e.value,e=>{h.value=G(e),h.value.forEach((o,_)=>{o.checked=!0}),S.value=G(h.value)},{deep:!0,immediate:!0});const Ee=e=>{C.value=e,E.value=!e.length},B=e=>{if(y.value&&C.value.length===0)w.$message.info("\u8BF7\u9009\u62E9\u6216\u67E5\u8BE2\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let o=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;switch(e.action){case"status":e.value===1?o=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`:o=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;break;case"delete":o=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;break;case"all-status":e.value===1?o=`\u662F\u5426\u786E\u5B9A\u5168\u90E8\u542F\u7528\u67E5\u8BE2\u7ED3\u679C\u7684${$.total}\u6761\u6570\u636E?`:o=`\u662F\u5426\u786E\u5B9A\u5168\u90E8\u7981\u7528\u67E5\u8BE2\u7ED3\u679C\u7684${$.total}\u6761\u6570\u636E?`;break;case"all-delete":o=`\u662F\u5426\u786E\u5B9A\u5168\u90E8\u5220\u9664\u67E5\u8BE2\u7ED3\u679C\u7684${$.total}\u6761\u6570\u636E?`;break}w.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:o,hideCancel:!1,onOk:()=>{p(!0),e.ids=C.value,zt({...e,...u.value}).then(_=>{p(!1),w.$message.success("\u64CD\u4F5C\u6210\u529F, \u4EFB\u52A1\u5DF2\u63D0\u4EA4"),V(),Z.value.selectAll(!1)})}})}},{copy:$e,copied:ae}=Dt(),we=async e=>{const{data:o}=await qt({id:e});$e(o.key)};de(ae,()=>{ae.value&&w.$message.success("\u590D\u5236\u6210\u529F")});const K=n(!1),q=n(),Fe=e=>{K.value=!0,q.value=e},Be=()=>{K.value=!1},P=n(!1),le=n(),De=e=>{P.value=!0,q.value=e,le.value="key"};return(e,o)=>{const _=je,v=at,g=lt,O=ot,L=ut,D=nt,f=st,oe=it,H=dt,Ve=rt,ue=ct,Ie=mt,r=pt,ne=Ge,se=_t,T=vt,Ae=Je,xe=ft,Se=yt,ze=Qe,qe=We,Le=gt,Ne=kt,Me=bt,Ue=ht,Ke=Ct,Pe=Et,Oe=$t,He=wt,Te=Ft;return c(),k("div",Ot,[a(g,{class:"container-breadcrumb"},{default:l(()=>[a(v,null,{default:l(()=>[a(_)]),_:1}),a(v,null,{default:l(()=>[s(d(e.$t("menu.key")),1)]),_:1}),a(v,null,{default:l(()=>[s(d(e.$t("menu.key.model.list")),1)]),_:1})]),_:1}),a(Te,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:l(()=>[a(H,null,{default:l(()=>[a(f,{flex:1},{default:l(()=>[a(Ve,{model:u.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:l(()=>[a(H,{gutter:16},{default:l(()=>[a(f,{span:8},{default:l(()=>[a(D,{field:"corp",label:e.$t("key.form.corp")},{default:l(()=>[a(L,{modelValue:u.value.corp,"onUpdate:modelValue":o[0]||(o[0]=t=>u.value.corp=t),placeholder:e.$t("key.form.selectDefault"),"allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),k(I,null,A(J.value,t=>(c(),N(O,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(D,{field:"key",label:e.$t("key.form.key")},{default:l(()=>[a(oe,{modelValue:u.value.key,"onUpdate:modelValue":o[1]||(o[1]=t=>u.value.key=t),placeholder:e.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(D,{field:"models",label:e.$t("key.form.models")},{default:l(()=>[a(L,{modelValue:u.value.models,"onUpdate:modelValue":o[2]||(o[2]=t=>u.value.models=t),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),k(I,null,A(Q.value,t=>(c(),N(O,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(D,{field:"model_agents",label:e.$t("key.form.modelAgents")},{default:l(()=>[a(L,{modelValue:u.value.model_agents,"onUpdate:modelValue":o[3]||(o[3]=t=>u.value.model_agents=t),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),k(I,null,A(W.value,t=>(c(),N(O,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(D,{field:"status",label:e.$t("key.form.status")},{default:l(()=>[a(L,{modelValue:u.value.status,"onUpdate:modelValue":o[4]||(o[4]=t=>u.value.status=t),options:M(ve),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(D,{field:"remark",label:e.$t("key.form.remark")},{default:l(()=>[a(oe,{modelValue:u.value.remark,"onUpdate:modelValue":o[5]||(o[5]=t=>u.value.remark=t),placeholder:e.$t("key.form.remark.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(ue,{style:{height:"84px"},direction:"vertical"}),a(f,{flex:"86px",style:{"text-align":"right"}},{default:l(()=>[a(se,{direction:"vertical",size:18},{default:l(()=>[a(r,{type:"primary",onClick:V},{icon:l(()=>[a(Ie)]),default:l(()=>[s(" "+d(e.$t("key.form.search")),1)]),_:1}),a(r,{onClick:ge},{icon:l(()=>[a(ne)]),default:l(()=>[s(" "+d(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(ue,{style:{"margin-top":"0","margin-bottom":"16px"}}),a(H,{style:{"margin-bottom":"16px"}},{default:l(()=>[a(f,{span:12},{default:l(()=>[a(se,null,{default:l(()=>[a(r,{type:"primary",onClick:o[6]||(o[6]=t=>e.$router.push({name:"KeyCreate"}))},{default:l(()=>[s(d(e.$t("key.operation.create")),1)]),_:1}),a(r,{type:"primary",status:"success",disabled:E.value,title:E.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[7]||(o[7]=t=>B({action:"status",value:1}))},{default:l(()=>[s(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),a(r,{type:"primary",status:"danger",disabled:E.value,title:E.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[8]||(o[8]=t=>B({action:"status",value:2}))},{default:l(()=>[s(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),a(r,{type:"primary",status:"danger",disabled:E.value,title:E.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[9]||(o[9]=t=>B({action:"delete"}))},{default:l(()=>[s(" \u5220\u9664 ")]),_:1},8,["disabled","title"]),a(r,{type:"primary",status:"success",disabled:y.value,title:y.value?"\u8BF7\u67E5\u8BE2\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[10]||(o[10]=t=>B({action:"all-status",value:1}))},{default:l(()=>[s(" \u5168\u90E8\u542F\u7528 ")]),_:1},8,["disabled","title"]),a(r,{type:"primary",status:"danger",disabled:y.value,title:y.value?"\u8BF7\u67E5\u8BE2\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[11]||(o[11]=t=>B({action:"all-status",value:2}))},{default:l(()=>[s(" \u5168\u90E8\u7981\u7528 ")]),_:1},8,["disabled","title"]),a(r,{type:"primary",status:"danger",disabled:y.value,title:y.value?"\u8BF7\u67E5\u8BE2\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[12]||(o[12]=t=>B({action:"all-delete"}))},{default:l(()=>[s(" \u5168\u90E8\u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),a(f,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:l(()=>[a(T,{content:e.$t("actions.refresh")},{default:l(()=>[b("div",{class:"action-icon",onClick:V},[a(ne,{size:"18"})])]),_:1},8,["content"]),a(Se,{onSelect:be},{content:l(()=>[(c(!0),k(I,null,A(M(pe),t=>(c(),N(xe,{key:t.value,value:t.value,class:Ze({active:t.value===U.value})},{default:l(()=>[b("span",null,d(t.name),1)]),_:2},1032,["value","class"]))),128))]),default:l(()=>[a(T,{content:e.$t("actions.density")},{default:l(()=>[b("div",Ht,[a(Ae,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(T,{content:e.$t("actions.column_setting")},{default:l(()=>[a(Ne,{trigger:"click",position:"bl",onPopupVisibleChange:Ce},{content:l(()=>[b("div",Rt,[(c(!0),k(I,null,A(S.value,(t,m)=>(c(),k("div",{key:t.dataIndex,class:"setting"},[b("div",jt,[a(qe)]),b("div",null,[a(Le,{modelValue:t.checked,"onUpdate:modelValue":R=>t.checked=R,onChange:R=>he(R,t,m)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),b("div",Gt,d(t.title==="#"?"\u5E8F\u5217\u53F7":t.title),1)]))),128))])]),default:l(()=>[b("div",Tt,[a(ze,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(Pe,{ref_key:"tableRef",ref:Z,"row-key":"id",loading:M(re),pagination:$,columns:h.value,data:Y.value,bordered:!1,size:U.value,"row-selection":ce,onPageChange:fe,onPageSizeChange:ye,onSelectionChange:Ee},{key:l(({record:t})=>[s(d(t.key.length>11?t.key.substr(0,10)+t.key.substr(-10):t.key)+" ",1),a(Me,{class:"copy-btn",onClick:m=>we(t.id)},null,8,["onClick"])]),model_names:l(({record:t})=>[t.model_names?(c(),k("span",Jt,[a(r,{type:"text",size:"small",onClick:m=>De(t.id)},{default:l(()=>[s(" \u67E5\u770B ")]),_:2},1032,["onClick"])])):(c(),k("span",Qt,d("-")))]),model_agent_names:l(({record:t})=>{var m;return[s(d(((m=t==null?void 0:t.model_agent_names)==null?void 0:m.join(","))||"-"),1)]}),used_quota:l(({record:t})=>[s(" $"+d(t.used_quota>0?M(It)(t.used_quota):"0.00"),1)]),remark:l(({record:t})=>[s(d(t.remark||"-"),1)]),status:l(({record:t})=>[a(Ue,{modelValue:t.status,"onUpdate:modelValue":m=>t.status=m,"checked-value":1,"unchecked-value":2,onChange:m=>ke({id:`${t.id}`,status:Number(`${t.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:l(({record:t})=>[a(r,{type:"text",size:"small",onClick:m=>Fe(t.id)},{default:l(()=>[s(d(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),a(r,{type:"text",size:"small",onClick:m=>e.$router.push({name:"KeyUpdate",query:{id:`${t.id}`}})},{default:l(()=>[s(d(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),a(Ke,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:m=>me({id:`${t.id}`})},{default:l(()=>[a(r,{type:"text",size:"small"},{default:l(()=>[s(d(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(Oe,{title:e.$t("menu.key.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:K.value,onCancel:Be},{default:l(()=>[a(Pt,{id:q.value},null,8,["id"])]),_:1},8,["title","visible"]),a(He,{visible:P.value,"onUpdate:visible":o[13]||(o[13]=t=>P.value=t),title:"\u6A21\u578B","modal-style":{padding:"25px 15px 20px 15px"},"unmount-on-close":"","hide-cancel":"",simple:"",width:"920px","ok-text":"\u5173\u95ED"},{default:l(()=>[a(Kt,{id:q.value,action:le.value},null,8,["id","action"])]),_:1},8,["visible"])]),_:1})])}}});const qa=Xe(Xt,[["__scopeId","data-v-b59dea33"]]);export{qa as default};
