import{u as Se,A as Ue,m as Le,v as Te,I as Be,w as Ne,_ as Ae}from"./index.74baba8a.js";import{u as Fe}from"./loading.dbaba456.js";/* empty css                *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{c as E,S as Re}from"./sortable.esm.507626e9.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              */import{d as Ee,r as ee,e as c,c as K,w as Ke,B as s,C as m,aH as l,aG as t,aL as f,aM as u,aD as w,aE as Me,aJ as q,aI as C,u as M,F as k,D as Pe,n as Oe,aK as je,aF as He,bC as Ge,bD as Je,b3 as Qe,bB as We,b2 as Xe,aT as Ye,bE as Ze,bF as xe,b6 as ea,bG as aa,ab as la,aV as ta,bj as oa,bk as na,bm as sa,bn as ua,b5 as ia,bH as da,bI as ra,bJ as pa,aU as ca,bK as ma,a$ as _a,bL as fa}from"./arco.d2aaf5b7.js";import{h as ya}from"./vue.ca65198a.js";import{s as va,q as ba}from"./key.2903752a.js";import{f as ka,b as ha}from"./app.097fee61.js";import{q as ga}from"./model.5554004c.js";import"./chart.61872c57.js";const $a={class:"container"},Va={class:"action-icon"},wa={class:"action-icon"},qa={id:"tableSetting"},Ca={style:{"margin-right":"4px",cursor:"move"}},Ia={class:"title"},za={key:0},Da={key:1},Sa={key:0},Ua={key:1},La={key:0,class:"circle"},Ta={key:1,class:"circle pass"},Ba={name:"AppKeyList"},Na=Ee({...Ba,setup(Aa){const ae=ya(),le=ee({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),P=c([]);(async()=>{try{const{data:e}=await ka();P.value=e.items}catch{}})();const S=c([]);(async()=>{try{const{data:e}=await ga();S.value=e.items}catch{}})();const te=async e=>{b(!0);try{await va(e),g()}catch{}finally{b(!1)}},O=()=>({type:1,app_id:c(),key:"",models:[],quota:c(),status:c(),created_at:[]}),{loading:oe,setLoading:b}=Fe(!0),{t:r}=Se(),j=c([]),i=c(O()),h=c([]),I=c([]),U=c("medium"),L={current:1,pageSize:10,showTotal:!0},T=ee({...L}),ne=K(()=>[{name:r("searchTable.size.mini"),value:"mini"},{name:r("searchTable.size.small"),value:"small"},{name:r("searchTable.size.medium"),value:"medium"},{name:r("searchTable.size.large"),value:"large"}]),se=K(()=>[{title:r("key.columns.app_id"),dataIndex:"app_id",slotName:"app_id"},{title:r("key.columns.key"),dataIndex:"key",slotName:"key"},{title:r("key.columns.quota"),dataIndex:"quota",slotName:"quota"},{title:r("key.columns.app.models"),dataIndex:"model_names",slotName:"model_names"},{title:r("key.columns.status"),dataIndex:"status",slotName:"status"},{title:r("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at"},{title:r("key.columns.operations"),dataIndex:"operations",slotName:"operations"}]),ue=K(()=>[{label:r("key.dict.status.1"),value:1},{label:r("key.dict.status.2"),value:2}]),g=async(e={current:1,pageSize:10,type:1,app_id:ae.query.app_id})=>{b(!0);try{const{data:o}=await ba(e);j.value=o.items,T.current=e.current,T.total=o.paging.total}catch{}finally{b(!1)}},H=()=>{g({...L,...i.value})},ie=e=>{g({...L,...i.value,current:e})};g();const de=()=>{i.value=O()},re=(e,o)=>{U.value=e},pe=(e,o,d)=>{e?h.value.splice(d,0,o):h.value=I.value.filter(_=>_.dataIndex!==o.dataIndex)},G=(e,o,d,_=!1)=>{const v=_?E(e):e;return o>-1&&d>-1&&v.splice(o,1,v.splice(d,1,v[o]).pop()),v},ce=e=>{e&&Oe(()=>{const o=document.getElementById("tableSetting");new Re(o,{onEnd(d){const{oldIndex:_,newIndex:v}=d;G(h.value,_,v),G(I.value,_,v)}})})};Ke(()=>se.value,e=>{h.value=E(e),h.value.forEach((o,d)=>{o.checked=!0}),I.value=E(h.value)},{deep:!0,immediate:!0});const $=c(!1),J=c(),n=c({}),me=async e=>{var o,d;b(!0);try{n.value.id=e.id,n.value.key=e.key,n.value.is_limit_quota=e.is_limit_quota,n.value.quota=e.quota,n.value.models=e.models,n.value.ip_whitelist=((o=e.ip_whitelist)==null?void 0:o.join(`
`))||"",n.value.ip_blacklist=((d=e.ip_blacklist)==null?void 0:d.join(`
`))||"",n.value.remark=e.remark,$.value=!0}catch{}finally{b(!1)}},_e=async e=>{var d;if(await((d=J.value)==null?void 0:d.validate())){$.value=!0,e(!1);return}b(!0);try{await ha(n.value),navigator.clipboard.writeText(n.value.key),e(),g()}catch{}finally{b(!1)}},fe=()=>{$.value=!1};return(e,o)=>{const d=Ue,_=je,v=He,B=Ge,z=Je,p=Qe,y=We,Q=Xe,W=Ye,ye=Ze,N=xe,X=ea,Y=aa,ve=la,V=ta,Z=Le,x=oa,A=na,be=Te,ke=sa,he=ua,ge=Be,$e=Ne,Ve=ia,we=da,qe=ra,Ce=pa,Ie=ca,F=ma,ze=_a,De=fa;return s(),m("div",$a,[l(v,{class:"container-breadcrumb"},{default:t(()=>[l(_,null,{default:t(()=>[l(d)]),_:1}),l(_,null,{default:t(()=>[f(u(e.$t("menu.key")),1)]),_:1}),l(_,null,{default:t(()=>[f(u(e.$t("menu.key.app.list")),1)]),_:1})]),_:1}),l(De,{class:"general-card",title:e.$t("menu.key.app.list"),bordered:!1},{extra:t(()=>[l(ze,{visible:$.value,"onUpdate:visible":o[13]||(o[13]=a=>$.value=a),title:e.$t("app.form.title.keyConfig"),"ok-text":e.$t("app.button.save"),onCancel:fe,onBeforeOk:_e},{default:t(()=>[l(X,{ref_key:"formRef",ref:J,model:n.value},{default:t(()=>[l(p,{field:"key",label:e.$t("app.label.key")},{default:t(()=>[l(Q,{modelValue:n.value.key,"onUpdate:modelValue":o[6]||(o[6]=a=>n.value.key=a),placeholder:e.$t("app.placeholder.key"),readonly:""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(p,{field:"is_limit_quota",label:e.$t("app.label.isLimitQuota")},{default:t(()=>[l(Ie,{modelValue:n.value.is_limit_quota,"onUpdate:modelValue":o[7]||(o[7]=a=>n.value.is_limit_quota=a)},null,8,["modelValue"])]),_:1},8,["label"]),n.value.is_limit_quota?(s(),w(p,{key:0,field:"quota",label:e.$t("app.label.quota"),rules:[{required:!0,message:e.$t("app.error.quota.required")}]},{default:t(()=>[l(W,{modelValue:n.value.quota,"onUpdate:modelValue":o[8]||(o[8]=a=>n.value.quota=a),placeholder:e.$t("app.placeholder.quota"),min:1},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):Me("",!0),l(p,{field:"models",label:e.$t("app.label.models")},{default:t(()=>[l(z,{modelValue:n.value.models,"onUpdate:modelValue":o[9]||(o[9]=a=>n.value.models=a),placeholder:e.$t("app.placeholder.models"),"max-tag-count":3,multiple:"","allow-clear":""},{default:t(()=>[(s(!0),m(q,null,C(S.value,a=>(s(),w(B,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),l(p,{field:"ip_whitelist",label:e.$t("app.label.ip_whitelist")},{default:t(()=>[l(F,{modelValue:n.value.ip_whitelist,"onUpdate:modelValue":o[10]||(o[10]=a=>n.value.ip_whitelist=a),placeholder:e.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(p,{field:"ip_blacklist",label:e.$t("app.label.ip_blacklist")},{default:t(()=>[l(F,{modelValue:n.value.ip_blacklist,"onUpdate:modelValue":o[11]||(o[11]=a=>n.value.ip_blacklist=a),placeholder:e.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:10}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(p,{field:"remark",label:e.$t("app.placeholder.remark")},{default:t(()=>[l(F,{modelValue:n.value.remark,"onUpdate:modelValue":o[12]||(o[12]=a=>n.value.remark=a),placeholder:e.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),default:t(()=>[l(N,null,{default:t(()=>[l(y,{flex:1},{default:t(()=>[l(X,{model:i.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[l(N,{gutter:16},{default:t(()=>[l(y,{span:8},{default:t(()=>[l(p,{field:"app_id",label:e.$t("key.form.app")},{default:t(()=>[l(z,{modelValue:i.value.app_id,"onUpdate:modelValue":o[0]||(o[0]=a=>i.value.app_id=a),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},{default:t(()=>[(s(!0),m(q,null,C(P.value,a=>(s(),w(B,{key:a.app_id,value:a.app_id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(p,{field:"key",label:e.$t("key.form.key")},{default:t(()=>[l(Q,{modelValue:i.value.key,"onUpdate:modelValue":o[1]||(o[1]=a=>i.value.key=a),placeholder:e.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(p,{field:"models",label:e.$t("key.form.models")},{default:t(()=>[l(z,{modelValue:i.value.models,"onUpdate:modelValue":o[2]||(o[2]=a=>i.value.models=a),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(s(!0),m(q,null,C(S.value,a=>(s(),w(B,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(p,{field:"quota",label:e.$t("key.form.quota")},{default:t(()=>[l(W,{modelValue:i.value.quota,"onUpdate:modelValue":o[3]||(o[3]=a=>i.value.quota=a),placeholder:e.$t("key.form.quota.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(p,{field:"status",label:e.$t("key.form.status")},{default:t(()=>[l(z,{modelValue:i.value.status,"onUpdate:modelValue":o[4]||(o[4]=a=>i.value.status=a),options:M(ue),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(p,{field:"created_at",label:e.$t("key.form.created_at")},{default:t(()=>[l(ye,{modelValue:i.value.created_at,"onUpdate:modelValue":o[5]||(o[5]=a=>i.value.created_at=a),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),l(Y,{style:{height:"84px"},direction:"vertical"}),l(y,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[l(x,{direction:"vertical",size:18},{default:t(()=>[l(V,{type:"primary",onClick:H},{icon:t(()=>[l(ve)]),default:t(()=>[f(" "+u(e.$t("key.form.search")),1)]),_:1}),l(V,{onClick:de},{icon:t(()=>[l(Z)]),default:t(()=>[f(" "+u(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),l(Y,{style:{"margin-top":"0"}}),l(N,{style:{"margin-bottom":"16px"}},{default:t(()=>[l(y,{span:12},{default:t(()=>[l(x)]),_:1}),l(y,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[l(A,{content:e.$t("searchTable.actions.refresh")},{default:t(()=>[k("div",{class:"action-icon",onClick:H},[l(Z,{size:"18"})])]),_:1},8,["content"]),l(he,{onSelect:re},{content:t(()=>[(s(!0),m(q,null,C(M(ne),a=>(s(),w(ke,{key:a.value,value:a.value,class:Pe({active:a.value===U.value})},{default:t(()=>[k("span",null,u(a.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[l(A,{content:e.$t("searchTable.actions.density")},{default:t(()=>[k("div",Va,[l(be,{size:"18"})])]),_:1},8,["content"])]),_:1}),l(A,{content:e.$t("searchTable.actions.columnSetting")},{default:t(()=>[l(we,{trigger:"click",position:"bl",onPopupVisibleChange:ce},{content:t(()=>[k("div",qa,[(s(!0),m(q,null,C(I.value,(a,D)=>(s(),m("div",{key:a.dataIndex,class:"setting"},[k("div",Ca,[l($e)]),k("div",null,[l(Ve,{modelValue:a.checked,"onUpdate:modelValue":R=>a.checked=R,onChange:R=>pe(R,a,D)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),k("div",Ia,u(a.title==="#"?"\u5E8F\u5217\u53F7":a.title),1)]))),128))])]),default:t(()=>[k("div",wa,[l(ge,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),l(Ce,{"row-key":"id",loading:M(oe),pagination:T,columns:h.value,data:j.value,bordered:!1,size:U.value,"row-selection":le,onPageChange:ie},{quota:t(({record:a})=>[a.is_limit_quota?(s(),m("span",za,u(a.quota),1)):(s(),m("span",Da,u(e.$t("key.columns.quota.no_limit")),1))]),model_names:t(({record:a})=>[a.model_names?(s(),m("span",Sa,u(a.model_names.join(",")),1)):(s(),m("span",Ua,u(e.$t("key.columns.app.models.no_limit")),1))]),type:t(({record:a})=>[f(u(e.$t(`key.dict.type.${a.type}`)),1)]),corp:t(({record:a})=>[f(u(e.$t(`key.dict.corp.${a.corp}`)),1)]),dataFormat:t(({record:a})=>[f(u(e.$t(`key.dict.data_format.${a.data_format}`)),1)]),status:t(({record:a})=>[a.status===3?(s(),m("span",La)):(s(),m("span",Ta)),f(" "+u(e.$t(`key.dict.status.${a.status}`)),1)]),operations:t(({record:a})=>[l(V,{type:"text",size:"small",onClick:D=>e.$router.push({name:"KeyDetail",query:{id:`${a.id}`}})},{default:t(()=>[f(u(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),l(V,{type:"text",size:"small",onClick:D=>me(a)},{default:t(()=>[f(u(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),l(qe,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:D=>te({id:`${a.id}`})},{default:t(()=>[l(V,{type:"text",size:"small"},{default:t(()=>[f(u(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const il=Ae(Na,[["__scopeId","data-v-7db6fc7d"]]);export{il as default};
