import{u as Ke,B as Pe,p as Oe,y as Te,i as He,z as Re,_ as je}from"./index.d6462cde.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as Ge,r as ne,e as u,c as T,w as se,B as c,C,aH as a,aG as l,aL as i,aM as d,aJ as I,aI as D,aD as A,u as x,F as v,D as Je,g as Qe,n as We,aK as Xe,aF as Ye,bA as Ze,bB as et,b2 as tt,bC as at,b1 as lt,bD as ot,b5 as nt,bE as st,ab as ut,aU as it,bi as dt,bj as rt,bl as ct,bm as mt,b4 as pt,bF as _t,ad as ft,aT as yt,bG as gt,bH as vt,aV as kt,bI as ht}from"./arco.54c7388d.js";import{h as bt,u as Ct}from"./vue.aa90ed69.js";import{u as wt}from"./loading.7321a6c2.js";import{q as $t}from"./common.df364eef.js";import{s as Vt,q as Ft,a as It,b as Dt,c as St}from"./key.432821f3.js";import{c as H,S as zt}from"./sortable.esm.777e758f.js";import{q as Bt}from"./model.e72c4173.js";import{f as Et}from"./agent.885aff6b.js";import{q as qt}from"./corp.e8a3355a.js";import{_ as At}from"./index.vue_vue_type_script_setup_true_lang.65e0ffe0.js";import"./chart.f14251fc.js";import"./base.87fcf6e2.js";/* empty css                *//* empty css                */const xt={class:"container"},Lt={class:"action-icon"},Nt={class:"action-icon"},Ut={id:"tableSetting"},Mt={style:{"margin-right":"4px",cursor:"move"}},Kt={class:"title"},Pt={name:"KeyList"},Ot=Ge({...Pt,setup(Tt){const{loading:ue,setLoading:m}=wt(!0),{proxy:w}=Qe(),S=bt(),ie=ne({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),R=u([]);(async()=>{m(!0);try{const{data:e}=await qt();R.value=e.items}catch{}finally{m(!1)}})();const j=u([]);(async()=>{try{const{data:e}=await Bt();j.value=e.items}catch{}})();const G=u([]);(async()=>{try{const{data:e}=await Et();G.value=e.items;const o=new Array(0);S.query.agent_id&&(o[0]=S.query.agent_id,n.value.model_agents=o)}catch{}})();const de=async e=>{m(!0);try{await Vt(e),w.$message.success("\u5220\u9664\u6210\u529F"),F()}catch{}finally{m(!1)}},J=()=>({type:2,corp:"",key:"",models:[],model_agents:u(),quota:u(),status:u(),remark:""}),{t:s}=Ke(),Q=u([]),n=u(J()),k=u([]),z=u([]),L=u("medium"),h=u([]),b=u(!0),W=u(),$={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},B=ne({...$}),re=T(()=>[{name:s("size.mini"),value:"mini"},{name:s("size.small"),value:"small"},{name:s("size.medium"),value:"medium"},{name:s("size.large"),value:"large"}]),ce=T(()=>[{title:s("key.columns.corp"),dataIndex:"corp_name",slotName:"corp_name",align:"center",width:110},{title:s("key.columns.key"),dataIndex:"key",slotName:"key",align:"center",width:230},{title:s("key.columns.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:s("key.columns.model_agents"),dataIndex:"model_agent_names",slotName:"model_agent_names",align:"center",ellipsis:!0,tooltip:!0},{title:s("key.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:s("common.weight"),dataIndex:"weight",slotName:"weight",align:"center",width:60},{title:s("key.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:s("key.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:s("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132},{title:s("key.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]),me=T(()=>[{label:s("key.dict.status.1"),value:1},{label:s("key.dict.status.2"),value:2}]),X=new Array(0);S.query.agent_id&&(X[0]=S.query.agent_id);const E=async(e={...$,type:2,model_agents:X})=>{m(!0);try{const{data:o}=await Ft(e);Q.value=o.items,B.current=e.current,B.pageSize=e.pageSize,B.total=o.paging.total}catch{}finally{m(!1)}},F=()=>{E({...$,...n.value})},pe=e=>{E({...$,...n.value,current:e})},_e=e=>{$.pageSize=e,E({...$,...n.value})};E();const fe=()=>{n.value=J()},ye=async e=>{m(!0);try{await It(e),w.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}catch{}finally{m(!1)}},ge=(e,o)=>{L.value=e},ve=(e,o,p)=>{e?k.value.splice(p,0,o):k.value=z.value.filter(_=>_.dataIndex!==o.dataIndex)},Y=(e,o,p,_=!1)=>{const y=_?H(e):e;return o>-1&&p>-1&&y.splice(o,1,y.splice(p,1,y[o]).pop()),y},ke=e=>{e&&We(()=>{const o=document.getElementById("tableSetting");new zt(o,{onEnd(p){const{oldIndex:_,newIndex:y}=p;Y(k.value,_,y),Y(z.value,_,y)}})})};se(()=>ce.value,e=>{k.value=H(e),k.value.forEach((o,p)=>{o.checked=!0}),z.value=H(k.value)},{deep:!0,immediate:!0});const he=e=>{h.value=e,b.value=!e.length},N=e=>{if(h.value.length===0)w.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let o=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;switch(e.action){case"status":e.value===1?o=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`:o=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;break;case"delete":o=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;break}w.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:o,hideCancel:!1,onOk:()=>{m(!0),e.ids=h.value,Dt(e).then(p=>{m(!1),w.$message.success("\u64CD\u4F5C\u6210\u529F"),F(),W.value.selectAll(!1)})}})}},{copy:be,copied:Z}=Ct(),Ce=async e=>{const{data:o}=await St({id:e});be(o.key)};se(Z,()=>{Z.value&&w.$message.success("\u590D\u5236\u6210\u529F")});const U=u(!1),ee=u(),we=e=>{U.value=!0,ee.value=e},$e=()=>{U.value=!1};return(e,o)=>{const p=Pe,_=Xe,y=Ye,M=Ze,q=et,V=tt,f=at,te=lt,K=ot,Ve=nt,ae=st,Fe=ut,g=it,le=Oe,oe=dt,P=rt,Ie=Te,De=ct,Se=mt,ze=He,Be=Re,Ee=pt,qe=_t,Ae=ft,xe=yt,Le=gt,Ne=vt,Ue=kt,Me=ht;return c(),C("div",xt,[a(y,{class:"container-breadcrumb"},{default:l(()=>[a(_,null,{default:l(()=>[a(p)]),_:1}),a(_,null,{default:l(()=>[i(d(e.$t("menu.key")),1)]),_:1}),a(_,null,{default:l(()=>[i(d(e.$t("menu.key.model.list")),1)]),_:1})]),_:1}),a(Me,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:l(()=>[a(K,null,{default:l(()=>[a(f,{flex:1},{default:l(()=>[a(Ve,{model:n.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:l(()=>[a(K,{gutter:16},{default:l(()=>[a(f,{span:8},{default:l(()=>[a(V,{field:"corp",label:e.$t("key.form.corp")},{default:l(()=>[a(q,{modelValue:n.value.corp,"onUpdate:modelValue":o[0]||(o[0]=t=>n.value.corp=t),placeholder:e.$t("key.form.selectDefault"),"allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),C(I,null,D(R.value,t=>(c(),A(M,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(V,{field:"key",label:e.$t("key.form.key")},{default:l(()=>[a(te,{modelValue:n.value.key,"onUpdate:modelValue":o[1]||(o[1]=t=>n.value.key=t),placeholder:e.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(V,{field:"models",label:e.$t("key.form.models")},{default:l(()=>[a(q,{modelValue:n.value.models,"onUpdate:modelValue":o[2]||(o[2]=t=>n.value.models=t),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),C(I,null,D(j.value,t=>(c(),A(M,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(V,{field:"model_agents",label:e.$t("key.form.modelAgents")},{default:l(()=>[a(q,{modelValue:n.value.model_agents,"onUpdate:modelValue":o[3]||(o[3]=t=>n.value.model_agents=t),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),C(I,null,D(G.value,t=>(c(),A(M,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(V,{field:"status",label:e.$t("key.form.status")},{default:l(()=>[a(q,{modelValue:n.value.status,"onUpdate:modelValue":o[4]||(o[4]=t=>n.value.status=t),options:x(me),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(f,{span:8},{default:l(()=>[a(V,{field:"remark",label:e.$t("key.form.remark")},{default:l(()=>[a(te,{modelValue:n.value.remark,"onUpdate:modelValue":o[5]||(o[5]=t=>n.value.remark=t),placeholder:e.$t("key.form.remark.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(ae,{style:{height:"84px"},direction:"vertical"}),a(f,{flex:"86px",style:{"text-align":"right"}},{default:l(()=>[a(oe,{direction:"vertical",size:18},{default:l(()=>[a(g,{type:"primary",onClick:F},{icon:l(()=>[a(Fe)]),default:l(()=>[i(" "+d(e.$t("key.form.search")),1)]),_:1}),a(g,{onClick:fe},{icon:l(()=>[a(le)]),default:l(()=>[i(" "+d(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(ae,{style:{"margin-top":"0","margin-bottom":"16px"}}),a(K,{style:{"margin-bottom":"16px"}},{default:l(()=>[a(f,{span:12},{default:l(()=>[a(oe,null,{default:l(()=>[a(g,{type:"primary",onClick:o[6]||(o[6]=t=>e.$router.push({name:"KeyCreate"}))},{default:l(()=>[i(d(e.$t("key.operation.create")),1)]),_:1}),a(g,{type:"primary",status:"success",disabled:b.value,title:b.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[7]||(o[7]=t=>N({action:"status",value:1}))},{default:l(()=>[i(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),a(g,{type:"primary",status:"danger",disabled:b.value,title:b.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[8]||(o[8]=t=>N({action:"status",value:2}))},{default:l(()=>[i(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),a(g,{type:"primary",status:"danger",disabled:b.value,title:b.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[9]||(o[9]=t=>N({action:"delete"}))},{default:l(()=>[i(" \u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),a(f,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:l(()=>[a(P,{content:e.$t("actions.refresh")},{default:l(()=>[v("div",{class:"action-icon",onClick:F},[a(le,{size:"18"})])]),_:1},8,["content"]),a(Se,{onSelect:ge},{content:l(()=>[(c(!0),C(I,null,D(x(re),t=>(c(),A(De,{key:t.value,value:t.value,class:Je({active:t.value===L.value})},{default:l(()=>[v("span",null,d(t.name),1)]),_:2},1032,["value","class"]))),128))]),default:l(()=>[a(P,{content:e.$t("actions.density")},{default:l(()=>[v("div",Lt,[a(Ie,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(P,{content:e.$t("actions.column_setting")},{default:l(()=>[a(qe,{trigger:"click",position:"bl",onPopupVisibleChange:ke},{content:l(()=>[v("div",Ut,[(c(!0),C(I,null,D(z.value,(t,r)=>(c(),C("div",{key:t.dataIndex,class:"setting"},[v("div",Mt,[a(Be)]),v("div",null,[a(Ee,{modelValue:t.checked,"onUpdate:modelValue":O=>t.checked=O,onChange:O=>ve(O,t,r)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),v("div",Kt,d(t.title==="#"?"\u5E8F\u5217\u53F7":t.title),1)]))),128))])]),default:l(()=>[v("div",Nt,[a(ze,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(Ne,{ref_key:"tableRef",ref:W,"row-key":"id",loading:x(ue),pagination:B,columns:k.value,data:Q.value,bordered:!1,size:L.value,"row-selection":ie,onPageChange:pe,onPageSizeChange:_e,onSelectionChange:he},{key:l(({record:t})=>[i(d(t.key.length>11?t.key.substr(0,10)+t.key.substr(-10):t.key)+" ",1),a(Ae,{class:"copy-btn",onClick:r=>Ce(t.id)},null,8,["onClick"])]),model_names:l(({record:t})=>{var r;return[i(d(((r=t==null?void 0:t.model_names)==null?void 0:r.join(","))||"-"),1)]}),model_agent_names:l(({record:t})=>{var r;return[i(d(((r=t==null?void 0:t.model_agent_names)==null?void 0:r.join(","))||"-"),1)]}),used_quota:l(({record:t})=>[i(" $"+d(t.used_quota>0?x($t)(t.used_quota):"0.00"),1)]),remark:l(({record:t})=>[i(d(t.remark||"-"),1)]),status:l(({record:t})=>[a(xe,{modelValue:t.status,"onUpdate:modelValue":r=>t.status=r,"checked-value":1,"unchecked-value":2,onChange:r=>ye({id:`${t.id}`,status:Number(`${t.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:l(({record:t})=>[a(g,{type:"text",size:"small",onClick:r=>we(t.id)},{default:l(()=>[i(d(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),a(g,{type:"text",size:"small",onClick:r=>e.$router.push({name:"KeyUpdate",query:{id:`${t.id}`}})},{default:l(()=>[i(d(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),a(Le,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:r=>de({id:`${t.id}`})},{default:l(()=>[a(g,{type:"text",size:"small"},{default:l(()=>[i(d(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(Ue,{title:e.$t("menu.key.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:U.value,onCancel:$e},{default:l(()=>[a(At,{id:ee.value},null,8,["id"])]),_:1},8,["title","visible"])]),_:1})])}}});const Ia=je(Ot,[["__scopeId","data-v-0d7a32fb"]]);export{Ia as default};
