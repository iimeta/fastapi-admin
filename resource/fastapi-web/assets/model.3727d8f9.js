import{u as ge,A as $e,m as we,v as Ve,I as Ce,w as Ie,_ as De}from"./index.74baba8a.js";import{u as ze}from"./loading.dbaba456.js";/* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{c as F,S as Se}from"./sortable.esm.507626e9.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              */import{d as Te,r as R,e as _,c as V,w as qe,B as m,C as v,aH as e,aG as a,aL as u,aM as s,u as C,aJ as L,aI as N,aD as j,F as f,D as Be,n as Fe,aK as Le,aF as Ne,bD as Ue,b3 as Ae,bB as Pe,b2 as Ee,bC as Ke,aT as Oe,bE as Me,bF as He,b6 as Re,bG as je,ab as xe,aV as Ge,bj as Je,a5 as Qe,bk as We,bm as Xe,bn as Ye,b5 as Ze,bH as et,bI as tt,bJ as at,bL as ot}from"./arco.d2aaf5b7.js";import{s as lt,q as nt}from"./key.2903752a.js";import{q as st}from"./model.5554004c.js";import"./chart.61872c57.js";import"./vue.ca65198a.js";const ct={class:"container"},ut={class:"action-icon"},rt={class:"action-icon"},it={id:"tableSetting"},dt={style:{"margin-right":"4px",cursor:"move"}},pt={class:"title"},mt={key:0,class:"circle"},_t={key:1,class:"circle pass"},ft={name:"KeyList"},yt=Te({...ft,setup(vt){const x=R({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),U=_([]);(async()=>{try{const{data:t}=await st();U.value=t.items}catch{}})();const G=async t=>{h(!0);try{await lt(t),$()}catch{}finally{h(!1)}},A=()=>({type:2,corp:"",key:"",models:[],quota:_(),status:_(),created_at:[]}),{loading:J,setLoading:h}=ze(!0),{t:c}=ge(),P=_([]),n=_(A()),y=_([]),g=_([]),I=_("medium"),D={current:1,pageSize:10,showTotal:!0},z=R({...D}),Q=V(()=>[{name:c("searchTable.size.mini"),value:"mini"},{name:c("searchTable.size.small"),value:"small"},{name:c("searchTable.size.medium"),value:"medium"},{name:c("searchTable.size.large"),value:"large"}]),W=V(()=>[{title:c("key.columns.corp"),dataIndex:"corp",slotName:"corp"},{title:c("key.columns.key"),dataIndex:"key",slotName:"key"},{title:c("key.columns.models"),dataIndex:"model_names",slotName:"model_names"},{title:c("key.columns.status"),dataIndex:"status",slotName:"status"},{title:c("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at"},{title:c("key.columns.operations"),dataIndex:"operations",slotName:"operations"}]),X=V(()=>[{label:c("key.dict.corp.OpenAI"),value:"OpenAI"}]),Y=V(()=>[{label:c("key.dict.status.1"),value:1},{label:c("key.dict.status.2"),value:2}]),$=async(t={current:1,pageSize:10,type:2})=>{h(!0);try{const{data:l}=await nt(t);P.value=l.items,z.current=t.current,z.total=l.paging.total}catch{}finally{h(!1)}},E=()=>{$({...D,...n.value})},Z=t=>{$({...D,...n.value,current:t})};$();const ee=()=>{n.value=A()},te=(t,l)=>{I.value=t},ae=(t,l,d)=>{t?y.value.splice(d,0,l):y.value=g.value.filter(r=>r.dataIndex!==l.dataIndex)},K=(t,l,d,r=!1)=>{const p=r?F(t):t;return l>-1&&d>-1&&p.splice(l,1,p.splice(d,1,p[l]).pop()),p},oe=t=>{t&&Fe(()=>{const l=document.getElementById("tableSetting");new Se(l,{onEnd(d){const{oldIndex:r,newIndex:p}=d;K(y.value,r,p),K(g.value,r,p)}})})};return qe(()=>W.value,t=>{y.value=F(t),y.value.forEach((l,d)=>{l.checked=!0}),g.value=F(y.value)},{deep:!0,immediate:!0}),(t,l)=>{const d=$e,r=Le,p=Ne,S=Ue,k=Ae,i=Pe,le=Ee,ne=Ke,se=Oe,ce=Me,T=He,ue=Re,O=je,re=xe,b=Ge,M=we,H=Je,ie=Qe,q=We,de=Ve,pe=Xe,me=Ye,_e=Ce,fe=Ie,ye=Ze,ve=et,ke=tt,be=at,he=ot;return m(),v("div",ct,[e(p,{class:"container-breadcrumb"},{default:a(()=>[e(r,null,{default:a(()=>[e(d)]),_:1}),e(r,null,{default:a(()=>[u(s(t.$t("menu.key")),1)]),_:1}),e(r,null,{default:a(()=>[u(s(t.$t("menu.key.model.list")),1)]),_:1})]),_:1}),e(he,{class:"general-card",title:t.$t("menu.key.model.list"),bordered:!1},{default:a(()=>[e(T,null,{default:a(()=>[e(i,{flex:1},{default:a(()=>[e(ue,{model:n.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(T,{gutter:16},{default:a(()=>[e(i,{span:8},{default:a(()=>[e(k,{field:"corp",label:t.$t("key.form.corp")},{default:a(()=>[e(S,{modelValue:n.value.corp,"onUpdate:modelValue":l[0]||(l[0]=o=>n.value.corp=o),options:C(X),placeholder:t.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:a(()=>[e(k,{field:"key",label:t.$t("key.form.key")},{default:a(()=>[e(le,{modelValue:n.value.key,"onUpdate:modelValue":l[1]||(l[1]=o=>n.value.key=o),placeholder:t.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:a(()=>[e(k,{field:"models",label:t.$t("key.form.models")},{default:a(()=>[e(S,{modelValue:n.value.models,"onUpdate:modelValue":l[2]||(l[2]=o=>n.value.models=o),placeholder:t.$t("key.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:a(()=>[(m(!0),v(L,null,N(U.value,o=>(m(),j(ne,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:a(()=>[e(k,{field:"quota",label:t.$t("key.form.quota")},{default:a(()=>[e(se,{modelValue:n.value.quota,"onUpdate:modelValue":l[3]||(l[3]=o=>n.value.quota=o),placeholder:t.$t("key.form.quota.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:a(()=>[e(k,{field:"status",label:t.$t("key.form.status")},{default:a(()=>[e(S,{modelValue:n.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>n.value.status=o),options:C(Y),placeholder:t.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:a(()=>[e(k,{field:"created_at",label:t.$t("key.form.created_at")},{default:a(()=>[e(ce,{modelValue:n.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=o=>n.value.created_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(O,{style:{height:"84px"},direction:"vertical"}),e(i,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(H,{direction:"vertical",size:18},{default:a(()=>[e(b,{type:"primary",onClick:E},{icon:a(()=>[e(re)]),default:a(()=>[u(" "+s(t.$t("key.form.search")),1)]),_:1}),e(b,{onClick:ee},{icon:a(()=>[e(M)]),default:a(()=>[u(" "+s(t.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(O,{style:{"margin-top":"0"}}),e(T,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(i,{span:12},{default:a(()=>[e(H,null,{default:a(()=>[e(b,{type:"primary",onClick:l[6]||(l[6]=o=>t.$router.push({name:"KeyCreate"}))},{icon:a(()=>[e(ie)]),default:a(()=>[u(" "+s(t.$t("key.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(i,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:a(()=>[e(q,{content:t.$t("searchTable.actions.refresh")},{default:a(()=>[f("div",{class:"action-icon",onClick:E},[e(M,{size:"18"})])]),_:1},8,["content"]),e(me,{onSelect:te},{content:a(()=>[(m(!0),v(L,null,N(C(Q),o=>(m(),j(pe,{key:o.value,value:o.value,class:Be({active:o.value===I.value})},{default:a(()=>[f("span",null,s(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(q,{content:t.$t("searchTable.actions.density")},{default:a(()=>[f("div",ut,[e(de,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(q,{content:t.$t("searchTable.actions.columnSetting")},{default:a(()=>[e(ve,{trigger:"click",position:"bl",onPopupVisibleChange:oe},{content:a(()=>[f("div",it,[(m(!0),v(L,null,N(g.value,(o,w)=>(m(),v("div",{key:o.dataIndex,class:"setting"},[f("div",dt,[e(fe)]),f("div",null,[e(ye,{modelValue:o.checked,"onUpdate:modelValue":B=>o.checked=B,onChange:B=>ae(B,o,w)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),f("div",pt,s(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:a(()=>[f("div",rt,[e(_e,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(be,{"row-key":"id",loading:C(J),pagination:z,columns:y.value,data:P.value,bordered:!1,size:I.value,"row-selection":x,onPageChange:Z},{type:a(({record:o})=>[u(s(t.$t(`key.dict.type.${o.type}`)),1)]),corp:a(({record:o})=>[u(s(t.$t(`key.dict.corp.${o.corp}`)),1)]),dataFormat:a(({record:o})=>[u(s(t.$t(`key.dict.data_format.${o.data_format}`)),1)]),status:a(({record:o})=>[o.status===3?(m(),v("span",mt)):(m(),v("span",_t)),u(" "+s(t.$t(`key.dict.status.${o.status}`)),1)]),operations:a(({record:o})=>[e(b,{type:"text",size:"small",onClick:w=>t.$router.push({name:"KeyDetail",query:{id:`${o.id}`}})},{default:a(()=>[u(s(t.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(b,{type:"text",size:"small",onClick:w=>t.$router.push({name:"KeyUpdate",query:{id:`${o.id}`}})},{default:a(()=>[u(s(t.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),e(ke,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:w=>G({id:`${o.id}`})},{default:a(()=>[e(b,{type:"text",size:"small"},{default:a(()=>[u(s(t.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const Kt=De(yt,[["__scopeId","data-v-5eeb7226"]]);export{Kt as default};
