import{u as he,G as ye,t as ge,B as $e,m as ke,C as we,_ as Ce}from"./index.3029f51b.js";import{u as Ve}from"./loading.fc5bfc3b.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css              */import{c as F,S as Ie}from"./sortable.esm.8e8df260.js";/* empty css               *//* empty css               *//* empty css              */import{d as ze,r as R,e as f,c as g,w as De,B as v,C as y,aH as e,aG as t,aL as c,aM as d,u as $,F as p,aJ as q,aI as G,aD as Se,D as Te,n as Ne,aK as Be,aF as Fe,bO as Me,b4 as Pe,bM as Ue,b3 as Oe,bP as xe,bQ as Ae,b7 as Le,bR as Ee,ab as Re,aW as qe,bn as Ge,a5 as He,bv as je,bw as Je,bx as Ke,b6 as Qe,bz as We,bS as Xe,bT as Ye,bV as Ze}from"./arco.fd20202f.js";import{s as et,a as tt}from"./model.943397e8.js";import"./chart.57980958.js";import"./vue.70a4bb93.js";const at={class:"container"},ot={class:"action-icon"},lt={class:"action-icon"},nt={id:"tableSetting"},st={style:{"margin-right":"4px",cursor:"move"}},dt={class:"title"},ct={key:0,class:"circle"},rt={key:1,class:"circle pass"},it={name:"ModelList"},ut=ze({...it,setup(mt){const H=R({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),j=async a=>{k(!0);try{await et(a),C()}catch{}finally{k(!1)}},M=()=>({corp:"",name:"",model:"",type:f(),status:f(),created_at:[]}),{loading:J,setLoading:k}=Ve(!0),{t:n}=he(),P=f([]),s=f(M()),_=f([]),w=f([]),I=f("medium"),z={current:1,pageSize:10,showTotal:!0},D=R({...z}),K=g(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),Q=g(()=>[{title:n("model.columns.corp"),dataIndex:"corp",slotName:"corp"},{title:n("model.columns.name"),dataIndex:"name",slotName:"name"},{title:n("model.columns.model"),dataIndex:"model",slotName:"model"},{title:n("model.columns.type"),dataIndex:"type",slotName:"type"},{title:n("model.columns.data_format"),dataIndex:"data_format",slotName:"dataFormat"},{title:n("model.columns.status"),dataIndex:"status",slotName:"status"},{title:n("model.columns.remark"),dataIndex:"remark",slotName:"remark"},{title:n("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at"},{title:n("model.columns.operations"),dataIndex:"operations",slotName:"operations"}]),W=g(()=>[{label:n("model.dict.corp.OpenAI"),value:"OpenAI"}]),X=g(()=>[{label:n("model.dict.type.1"),value:1},{label:n("model.dict.type.2"),value:2},{label:n("model.dict.type.3"),value:3},{label:n("model.dict.type.4"),value:4}]),Y=g(()=>[{label:n("model.dict.status.1"),value:1},{label:n("model.dict.status.2"),value:2}]),C=async(a={current:1,pageSize:10})=>{k(!0);try{const{data:l}=await tt(a);P.value=l.items,D.current=a.current,D.total=l.paging.total}catch{}finally{k(!1)}},U=()=>{C({...z,...s.value})},Z=a=>{C({...z,...s.value,current:a})};C();const ee=()=>{s.value=M()},te=(a,l)=>{I.value=a},ae=(a,l,u)=>{a?_.value.splice(u,0,l):_.value=w.value.filter(r=>r.dataIndex!==l.dataIndex)},O=(a,l,u,r=!1)=>{const m=r?F(a):a;return l>-1&&u>-1&&m.splice(l,1,m.splice(u,1,m[l]).pop()),m},oe=a=>{a&&Ne(()=>{const l=document.getElementById("tableSetting");new Ie(l,{onEnd(u){const{oldIndex:r,newIndex:m}=u;O(_.value,r,m),O(w.value,r,m)}})})};return De(()=>Q.value,a=>{_.value=F(a),_.value.forEach((l,u)=>{l.checked=!0}),w.value=F(_.value)},{deep:!0,immediate:!0}),(a,l)=>{const u=ye,r=Be,m=Fe,S=Me,b=Pe,i=Ue,x=Oe,le=xe,T=Ae,ne=Le,A=Ee,se=Re,h=qe,L=ge,E=Ge,de=He,N=je,ce=$e,re=Je,ie=Ke,ue=ke,me=we,pe=Qe,_e=We,fe=Xe,ve=Ye,be=Ze;return v(),y("div",at,[e(m,{class:"container-breadcrumb"},{default:t(()=>[e(r,null,{default:t(()=>[e(u)]),_:1}),e(r,null,{default:t(()=>[c(d(a.$t("menu.model")),1)]),_:1}),e(r,null,{default:t(()=>[c(d(a.$t("menu.model.list")),1)]),_:1})]),_:1}),e(be,{class:"general-card",title:a.$t("menu.model.list")},{default:t(()=>[e(T,null,{default:t(()=>[e(i,{flex:1},{default:t(()=>[e(ne,{model:s.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[e(T,{gutter:16},{default:t(()=>[e(i,{span:8},{default:t(()=>[e(b,{field:"corp",label:a.$t("model.form.corp")},{default:t(()=>[e(S,{modelValue:s.value.corp,"onUpdate:modelValue":l[0]||(l[0]=o=>s.value.corp=o),options:$(W),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:t(()=>[e(b,{field:"name",label:a.$t("model.form.name")},{default:t(()=>[e(x,{modelValue:s.value.name,"onUpdate:modelValue":l[1]||(l[1]=o=>s.value.name=o),placeholder:a.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:t(()=>[e(b,{field:"model",label:a.$t("model.form.model")},{default:t(()=>[e(x,{modelValue:s.value.model,"onUpdate:modelValue":l[2]||(l[2]=o=>s.value.model=o),placeholder:a.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:t(()=>[e(b,{field:"type",label:a.$t("model.form.type")},{default:t(()=>[e(S,{modelValue:s.value.type,"onUpdate:modelValue":l[3]||(l[3]=o=>s.value.type=o),options:$(X),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:t(()=>[e(b,{field:"status",label:a.$t("model.form.status")},{default:t(()=>[e(S,{modelValue:s.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>s.value.status=o),options:$(Y),placeholder:a.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(i,{span:8},{default:t(()=>[e(b,{field:"created_at",label:a.$t("model.form.created_at")},{default:t(()=>[e(le,{modelValue:s.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=o=>s.value.created_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(A,{style:{height:"84px"},direction:"vertical"}),e(i,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[e(E,{direction:"vertical",size:18},{default:t(()=>[e(h,{type:"primary",onClick:U},{icon:t(()=>[e(se)]),default:t(()=>[c(" "+d(a.$t("model.form.search")),1)]),_:1}),e(h,{onClick:ee},{icon:t(()=>[e(L)]),default:t(()=>[c(" "+d(a.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(A,{style:{"margin-top":"0"}}),e(T,{style:{"margin-bottom":"16px"}},{default:t(()=>[e(i,{span:12},{default:t(()=>[e(E,null,{default:t(()=>[e(h,{type:"primary",onClick:l[6]||(l[6]=o=>a.$router.push({name:"ModelCreate"}))},{icon:t(()=>[e(de)]),default:t(()=>[c(" "+d(a.$t("model.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(i,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[e(N,{content:a.$t("searchTable.actions.refresh")},{default:t(()=>[p("div",{class:"action-icon",onClick:U},[e(L,{size:"18"})])]),_:1},8,["content"]),e(ie,{onSelect:te},{content:t(()=>[(v(!0),y(q,null,G($(K),o=>(v(),Se(re,{key:o.value,value:o.value,class:Te({active:o.value===I.value})},{default:t(()=>[p("span",null,d(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[e(N,{content:a.$t("searchTable.actions.density")},{default:t(()=>[p("div",ot,[e(ce,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(N,{content:a.$t("searchTable.actions.columnSetting")},{default:t(()=>[e(_e,{trigger:"click",position:"bl",onPopupVisibleChange:oe},{content:t(()=>[p("div",nt,[(v(!0),y(q,null,G(w.value,(o,V)=>(v(),y("div",{key:o.dataIndex,class:"setting"},[p("div",st,[e(me)]),p("div",null,[e(pe,{modelValue:o.checked,"onUpdate:modelValue":B=>o.checked=B,onChange:B=>ae(B,o,V)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),p("div",dt,d(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:t(()=>[p("div",lt,[e(ue,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(ve,{"row-key":"id",loading:$(J),pagination:D,columns:_.value,data:P.value,bordered:!1,size:I.value,"row-selection":H,onPageChange:Z},{type:t(({record:o})=>[c(d(a.$t(`model.dict.type.${o.type}`)),1)]),corp:t(({record:o})=>[c(d(a.$t(`model.dict.corp.${o.corp}`)),1)]),dataFormat:t(({record:o})=>[c(d(a.$t(`model.dict.data_format.${o.data_format}`)),1)]),status:t(({record:o})=>[o.status===3?(v(),y("span",ct)):(v(),y("span",rt)),c(" "+d(a.$t(`model.dict.status.${o.status}`)),1)]),operations:t(({record:o})=>[e(h,{type:"text",size:"small",onClick:V=>a.$router.push({name:"ModelDetail",query:{id:`${o.id}`}})},{default:t(()=>[c(d(a.$t("model.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(h,{type:"text",size:"small",onClick:V=>a.$router.push({name:"ModelUpdate",query:{id:`${o.id}`}})},{default:t(()=>[c(d(a.$t("model.columns.operations.update")),1)]),_:2},1032,["onClick"]),e(fe,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:V=>j({id:`${o.id}`})},{default:t(()=>[e(h,{type:"text",size:"small"},{default:t(()=>[c(d(a.$t("model.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const St=Ce(ut,[["__scopeId","data-v-639eae35"]]);export{St as default};