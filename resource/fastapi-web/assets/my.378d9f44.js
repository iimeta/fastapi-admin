import{u as be,E as ge,o as ve,x as he,I as ye,y as we,_ as Ie}from"./index.525e8576.js";/* empty css               *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css              */import{c as D,S as Ve}from"./sortable.esm.2109e0e3.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as $e,r as E,e as f,c as y,w as Ce,B as b,C as h,aH as e,aG as t,aL as c,aM as d,u as w,F as p,aJ as q,aI as O,aD as ke,D as xe,n as Se,aK as ze,aF as De,bC as Ne,b2 as Te,bA as Fe,b1 as Be,bD as Le,bE as Me,b5 as Ue,bF as Pe,ab as Ae,aU as Ee,bi as qe,bj as Oe,bl as je,bm as Ge,b4 as Re,bG as He,bI as Je,bJ as Ke}from"./arco.aed15247.js";import{u as Qe}from"./loading.b5911e1d.js";import{a as We}from"./model.570552fc.js";import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const Xe={class:"container"},Ye={class:"action-icon"},Ze={class:"action-icon"},et={id:"tableSetting"},tt={style:{"margin-right":"4px",cursor:"move"}},at={class:"title"},ot={key:0,class:"circle red"},lt={key:1,class:"circle"},nt={name:"ModelList"},st=$e({...nt,setup(dt){const j=E({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),N=()=>({corp:"",name:"",model:"",type:f(),status:f(),created_at:[]}),{loading:G,setLoading:T}=Qe(!0),{t:n}=be(),F=f([]),s=f(N()),_=f([]),I=f([]),C=f("medium"),g={current:1,pageSize:10,showTotal:!0,showPageSize:!0},V=E({...g}),R=y(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),H=y(()=>[{title:n("model.columns.corp"),dataIndex:"corp",slotName:"corp",align:"center",width:110},{title:n("model.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:n("model.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:n("model.columns.type"),dataIndex:"type",slotName:"type",align:"center"},{title:n("model.columns.prompt_ratio"),dataIndex:"prompt_ratio",slotName:"prompt_ratio",align:"center"},{title:n("model.columns.completion_ratio"),dataIndex:"completion_ratio",slotName:"completion_ratio",align:"center"},{title:n("model.columns.fixed_quota"),dataIndex:"fixed_quota",slotName:"fixed_quota",align:"center"},{title:n("model.columns.data_format"),dataIndex:"data_format",slotName:"dataFormat",align:"center"},{title:n("model.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:80},{title:n("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:170}]),J=y(()=>[{label:n("model.dict.corp.OpenAI"),value:"OpenAI"},{label:n("model.dict.corp.Midjourney"),value:"Midjourney"},{label:n("model.dict.corp.GLM"),value:"GLM"}]),K=y(()=>[{label:n("model.dict.type.1"),value:1},{label:n("model.dict.type.2"),value:2},{label:n("model.dict.type.3"),value:3},{label:n("model.dict.type.4"),value:4}]),Q=y(()=>[{label:n("model.dict.status.1"),value:1},{label:n("model.dict.status.2"),value:2}]),$=async(a={...g})=>{T(!0);try{const{data:l}=await We(a);F.value=l.items,V.current=a.current,V.pageSize=a.pageSize,V.total=l.paging.total}catch{}finally{T(!1)}},B=()=>{$({...g,...s.value})},W=a=>{$({...g,...s.value,current:a})},X=a=>{g.pageSize=a,$({...g,...s.value})};$();const Y=()=>{s.value=N()},Z=(a,l)=>{C.value=a},ee=(a,l,i)=>{a?_.value.splice(i,0,l):_.value=I.value.filter(r=>r.dataIndex!==l.dataIndex)},L=(a,l,i,r=!1)=>{const m=r?D(a):a;return l>-1&&i>-1&&m.splice(l,1,m.splice(i,1,m[l]).pop()),m},te=a=>{a&&Se(()=>{const l=document.getElementById("tableSetting");new Ve(l,{onEnd(i){const{oldIndex:r,newIndex:m}=i;L(_.value,r,m),L(I.value,r,m)}})})};return Ce(()=>H.value,a=>{_.value=D(a),_.value.forEach((l,i)=>{l.checked=!0}),I.value=D(_.value)},{deep:!0,immediate:!0}),(a,l)=>{const i=ge,r=ze,m=De,k=Ne,v=Te,u=Fe,M=Be,ae=Le,x=Me,oe=Ue,U=Pe,le=Ae,P=Ee,A=ve,ne=qe,S=Oe,se=he,de=je,ce=Ge,ie=ye,re=we,me=Re,ue=He,pe=Je,_e=Ke;return b(),h("div",Xe,[e(m,{class:"container-breadcrumb"},{default:t(()=>[e(r,null,{default:t(()=>[e(i)]),_:1}),e(r,null,{default:t(()=>[c(d(a.$t("menu.my.model")),1)]),_:1})]),_:1}),e(_e,{class:"general-card",title:a.$t("menu.model.list"),bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"0 20px 20px"}},{default:t(()=>[e(x,null,{default:t(()=>[e(u,{flex:1},{default:t(()=>[e(oe,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[e(x,{gutter:16},{default:t(()=>[e(u,{span:8},{default:t(()=>[e(v,{field:"corp",label:a.$t("model.form.corp")},{default:t(()=>[e(k,{modelValue:s.value.corp,"onUpdate:modelValue":l[0]||(l[0]=o=>s.value.corp=o),options:w(J),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:t(()=>[e(v,{field:"name",label:a.$t("model.form.name")},{default:t(()=>[e(M,{modelValue:s.value.name,"onUpdate:modelValue":l[1]||(l[1]=o=>s.value.name=o),placeholder:a.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:t(()=>[e(v,{field:"model",label:a.$t("model.form.model")},{default:t(()=>[e(M,{modelValue:s.value.model,"onUpdate:modelValue":l[2]||(l[2]=o=>s.value.model=o),placeholder:a.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:t(()=>[e(v,{field:"type",label:a.$t("model.form.type")},{default:t(()=>[e(k,{modelValue:s.value.type,"onUpdate:modelValue":l[3]||(l[3]=o=>s.value.type=o),options:w(K),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:t(()=>[e(v,{field:"status",label:a.$t("model.form.status")},{default:t(()=>[e(k,{modelValue:s.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>s.value.status=o),options:w(Q),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:t(()=>[e(v,{field:"created_at",label:a.$t("model.form.created_at")},{default:t(()=>[e(ae,{modelValue:s.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=o=>s.value.created_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(U,{style:{height:"84px"},direction:"vertical"}),e(u,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[e(ne,{direction:"vertical",size:18},{default:t(()=>[e(P,{type:"primary",onClick:B},{icon:t(()=>[e(le)]),default:t(()=>[c(" "+d(a.$t("model.form.search")),1)]),_:1}),e(P,{onClick:Y},{icon:t(()=>[e(A)]),default:t(()=>[c(" "+d(a.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(U,{style:{"margin-top":"0"}}),e(x,{style:{"margin-bottom":"16px"}},{default:t(()=>[e(u,{span:24,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[e(S,{content:a.$t("searchTable.actions.refresh")},{default:t(()=>[p("div",{class:"action-icon",onClick:B},[e(A,{size:"18"})])]),_:1},8,["content"]),e(ce,{onSelect:Z},{content:t(()=>[(b(!0),h(q,null,O(w(R),o=>(b(),ke(de,{key:o.value,value:o.value,class:xe({active:o.value===C.value})},{default:t(()=>[p("span",null,d(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[e(S,{content:a.$t("searchTable.actions.density")},{default:t(()=>[p("div",Ye,[e(se,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(S,{content:a.$t("searchTable.actions.columnSetting")},{default:t(()=>[e(ue,{trigger:"click",position:"bl",onPopupVisibleChange:te},{content:t(()=>[p("div",et,[(b(!0),h(q,null,O(I.value,(o,fe)=>(b(),h("div",{key:o.dataIndex,class:"setting"},[p("div",tt,[e(re)]),p("div",null,[e(me,{modelValue:o.checked,"onUpdate:modelValue":z=>o.checked=z,onChange:z=>ee(z,o,fe)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),p("div",at,d(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:t(()=>[p("div",Ze,[e(ie,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(pe,{"row-key":"id",loading:w(G),pagination:V,columns:_.value,data:F.value,bordered:!1,size:C.value,"row-selection":j,onPageChange:W,onPageSizeChange:X},{type:t(({record:o})=>[c(d(a.$t(`model.dict.type.${o.type}`)),1)]),corp:t(({record:o})=>[c(d(a.$t(`model.dict.corp.${o.corp}`)),1)]),dataFormat:t(({record:o})=>[c(d(a.$t(`model.dict.data_format.${o.data_format}`)),1)]),prompt_ratio:t(({record:o})=>[c(d(o.billing_method===1?o.prompt_ratio:"-"),1)]),completion_ratio:t(({record:o})=>[c(d(o.billing_method===1?o.completion_ratio:"-"),1)]),fixed_quota:t(({record:o})=>[c(d(o.billing_method===2?o.fixed_quota:"-"),1)]),status:t(({record:o})=>[o.status===2?(b(),h("span",ot)):(b(),h("span",lt)),c(" "+d(a.$t(`model.dict.status.${o.status}`)),1)]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const zt=Ie(st,[["__scopeId","data-v-81e5db42"]]);export{zt as default};
