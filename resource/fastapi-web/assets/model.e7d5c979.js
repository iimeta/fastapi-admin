import{u as ge,F as $e,t as we,B as Ve,m as Ce,C as Ie,_ as ze}from"./index.3029f51b.js";import{u as De}from"./loading.fc5bfc3b.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css              */import{c as B,S as Se}from"./sortable.esm.8e8df260.js";/* empty css               *//* empty css               *//* empty css              */import{d as qe,r as R,e as _,c as V,w as Te,B as m,C as k,aH as e,aG as a,aL as u,aM as c,u as C,aJ as F,aI as L,aD as H,F as f,D as Ne,n as Be,aK as Fe,aF as Le,bO as Pe,b4 as Ue,bM as Oe,b3 as Ae,bN as Ke,aT as Me,bP as Ee,bQ as xe,b7 as Re,bR as He,ab as je,aW as Ge,bn as Je,a5 as Qe,bv as We,bw as Xe,bx as Ye,b6 as Ze,bz as et,bS as tt,bT as at,bV as ot}from"./arco.fd20202f.js";import{s as lt,q as nt}from"./key.3d211d3d.js";import{q as st}from"./model.943397e8.js";import"./chart.57980958.js";import"./vue.70a4bb93.js";const ct={class:"container"},ut={class:"action-icon"},rt={class:"action-icon"},dt={id:"tableSetting"},it={style:{"margin-right":"4px",cursor:"move"}},pt={class:"title"},mt={key:0,class:"circle"},_t={key:1,class:"circle pass"},ft={name:"KeyList"},yt=qe({...ft,setup(kt){const j=R({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),P=_([]);(async()=>{try{const{data:t}=await st();P.value=t.items}catch{}})();const G=async t=>{h(!0);try{await lt(t),$()}catch{}finally{h(!1)}},U=()=>({type:2,corp:"",key:"",models:[],quota:_(),status:_(),created_at:[]}),{loading:J,setLoading:h}=De(!0),{t:s}=ge(),O=_([]),n=_(U()),y=_([]),g=_([]),I=_("medium"),z={current:1,pageSize:10,showTotal:!0},D=R({...z}),Q=V(()=>[{name:s("searchTable.size.mini"),value:"mini"},{name:s("searchTable.size.small"),value:"small"},{name:s("searchTable.size.medium"),value:"medium"},{name:s("searchTable.size.large"),value:"large"}]),W=V(()=>[{title:s("key.columns.corp"),dataIndex:"corp",slotName:"corp"},{title:s("key.columns.key"),dataIndex:"key",slotName:"key"},{title:s("key.columns.quota"),dataIndex:"quota",slotName:"quota"},{title:s("key.columns.models"),dataIndex:"model_names",slotName:"model_names"},{title:s("key.columns.status"),dataIndex:"status",slotName:"status"},{title:s("key.columns.remark"),dataIndex:"remark",slotName:"remark"},{title:s("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at"},{title:s("key.columns.operations"),dataIndex:"operations",slotName:"operations"}]),X=V(()=>[{label:s("key.dict.corp.OpenAI"),value:"OpenAI"}]),Y=V(()=>[{label:s("key.dict.status.1"),value:1},{label:s("key.dict.status.2"),value:2}]),$=async(t={current:1,pageSize:10,type:2})=>{h(!0);try{const{data:l}=await nt(t);O.value=l.items,D.current=t.current,D.total=l.paging.total}catch{}finally{h(!1)}},A=()=>{$({...z,...n.value})},Z=t=>{$({...z,...n.value,current:t})};$();const ee=()=>{n.value=U()},te=(t,l)=>{I.value=t},ae=(t,l,i)=>{t?y.value.splice(i,0,l):y.value=g.value.filter(r=>r.dataIndex!==l.dataIndex)},K=(t,l,i,r=!1)=>{const p=r?B(t):t;return l>-1&&i>-1&&p.splice(l,1,p.splice(i,1,p[l]).pop()),p},oe=t=>{t&&Be(()=>{const l=document.getElementById("tableSetting");new Se(l,{onEnd(i){const{oldIndex:r,newIndex:p}=i;K(y.value,r,p),K(g.value,r,p)}})})};return Te(()=>W.value,t=>{y.value=B(t),y.value.forEach((l,i)=>{l.checked=!0}),g.value=B(y.value)},{deep:!0,immediate:!0}),(t,l)=>{const i=$e,r=Fe,p=Le,S=Pe,v=Ue,d=Oe,le=Ae,ne=Ke,se=Me,ce=Ee,q=xe,ue=Re,M=He,re=je,b=Ge,E=we,x=Je,de=Qe,T=We,ie=Ve,pe=Xe,me=Ye,_e=Ce,fe=Ie,ye=Ze,ke=et,ve=tt,be=at,he=ot;return m(),k("div",ct,[e(p,{class:"container-breadcrumb"},{default:a(()=>[e(r,null,{default:a(()=>[e(i)]),_:1}),e(r,null,{default:a(()=>[u(c(t.$t("menu.key")),1)]),_:1}),e(r,null,{default:a(()=>[u(c(t.$t("menu.key.model.list")),1)]),_:1})]),_:1}),e(he,{class:"general-card",title:t.$t("menu.key.model.list")},{default:a(()=>[e(q,null,{default:a(()=>[e(d,{flex:1},{default:a(()=>[e(ue,{model:n.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(q,{gutter:16},{default:a(()=>[e(d,{span:8},{default:a(()=>[e(v,{field:"corp",label:t.$t("key.form.corp")},{default:a(()=>[e(S,{modelValue:n.value.corp,"onUpdate:modelValue":l[0]||(l[0]=o=>n.value.corp=o),options:C(X),placeholder:t.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(v,{field:"key",label:t.$t("key.form.key")},{default:a(()=>[e(le,{modelValue:n.value.key,"onUpdate:modelValue":l[1]||(l[1]=o=>n.value.key=o),placeholder:t.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(v,{field:"models",label:t.$t("key.form.models")},{default:a(()=>[e(S,{modelValue:n.value.models,"onUpdate:modelValue":l[2]||(l[2]=o=>n.value.models=o),placeholder:t.$t("key.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:a(()=>[(m(!0),k(F,null,L(P.value,o=>(m(),H(ne,{key:o.id,value:o.id,label:o.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(v,{field:"quota",label:t.$t("key.form.quota")},{default:a(()=>[e(se,{modelValue:n.value.quota,"onUpdate:modelValue":l[3]||(l[3]=o=>n.value.quota=o),placeholder:t.$t("key.form.quota.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(v,{field:"status",label:t.$t("key.form.status")},{default:a(()=>[e(S,{modelValue:n.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>n.value.status=o),options:C(Y),placeholder:t.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(v,{field:"created_at",label:t.$t("key.form.created_at")},{default:a(()=>[e(ce,{modelValue:n.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=o=>n.value.created_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(M,{style:{height:"84px"},direction:"vertical"}),e(d,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(x,{direction:"vertical",size:18},{default:a(()=>[e(b,{type:"primary",onClick:A},{icon:a(()=>[e(re)]),default:a(()=>[u(" "+c(t.$t("key.form.search")),1)]),_:1}),e(b,{onClick:ee},{icon:a(()=>[e(E)]),default:a(()=>[u(" "+c(t.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(M,{style:{"margin-top":"0"}}),e(q,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(d,{span:12},{default:a(()=>[e(x,null,{default:a(()=>[e(b,{type:"primary",onClick:l[6]||(l[6]=o=>t.$router.push({name:"KeyCreate"}))},{icon:a(()=>[e(de)]),default:a(()=>[u(" "+c(t.$t("key.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(d,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:a(()=>[e(T,{content:t.$t("searchTable.actions.refresh")},{default:a(()=>[f("div",{class:"action-icon",onClick:A},[e(E,{size:"18"})])]),_:1},8,["content"]),e(me,{onSelect:te},{content:a(()=>[(m(!0),k(F,null,L(C(Q),o=>(m(),H(pe,{key:o.value,value:o.value,class:Ne({active:o.value===I.value})},{default:a(()=>[f("span",null,c(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(T,{content:t.$t("searchTable.actions.density")},{default:a(()=>[f("div",ut,[e(ie,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(T,{content:t.$t("searchTable.actions.columnSetting")},{default:a(()=>[e(ke,{trigger:"click",position:"bl",onPopupVisibleChange:oe},{content:a(()=>[f("div",dt,[(m(!0),k(F,null,L(g.value,(o,w)=>(m(),k("div",{key:o.dataIndex,class:"setting"},[f("div",it,[e(fe)]),f("div",null,[e(ye,{modelValue:o.checked,"onUpdate:modelValue":N=>o.checked=N,onChange:N=>ae(N,o,w)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),f("div",pt,c(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:a(()=>[f("div",rt,[e(_e,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(be,{"row-key":"id",loading:C(J),pagination:D,columns:y.value,data:O.value,bordered:!1,size:I.value,"row-selection":j,onPageChange:Z},{type:a(({record:o})=>[u(c(t.$t(`key.dict.type.${o.type}`)),1)]),corp:a(({record:o})=>[u(c(t.$t(`key.dict.corp.${o.corp}`)),1)]),dataFormat:a(({record:o})=>[u(c(t.$t(`key.dict.data_format.${o.data_format}`)),1)]),status:a(({record:o})=>[o.status===3?(m(),k("span",mt)):(m(),k("span",_t)),u(" "+c(t.$t(`key.dict.status.${o.status}`)),1)]),operations:a(({record:o})=>[e(b,{type:"text",size:"small",onClick:w=>t.$router.push({name:"KeyDetail",query:{id:`${o.id}`}})},{default:a(()=>[u(c(t.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(b,{type:"text",size:"small",onClick:w=>t.$router.push({name:"KeyUpdate",query:{id:`${o.id}`}})},{default:a(()=>[u(c(t.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),e(ve,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:w=>G({id:`${o.id}`})},{default:a(()=>[e(b,{type:"text",size:"small"},{default:a(()=>[u(c(t.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const Pt=ze(yt,[["__scopeId","data-v-8d0191f2"]]);export{Pt as default};
