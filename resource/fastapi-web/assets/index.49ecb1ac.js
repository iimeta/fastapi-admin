import{u as Ve,C as Ce,E as Ie,F as ze,G as De,m as Ue,v as Se,I as qe,w as Be,_ as Te}from"./index.74baba8a.js";import{u as Fe}from"./loading.dbaba456.js";/* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{c as T,S as Ne}from"./sortable.esm.507626e9.js";/* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css              */import{d as Ee,r as H,e as d,c as F,w as Le,B as h,C as y,aH as e,aG as t,aL as p,aM as i,u as N,F as f,aJ as j,aI as J,aD as Pe,D as xe,n as Ae,aX as Me,aK as Oe,aF as Qe,aT as Re,b3 as Ge,bB as He,b2 as je,bD as Je,bE as Ke,bF as Xe,b6 as We,bG as Ye,ab as Ze,aV as ea,bj as aa,a5 as ta,bk as la,bm as oa,bn as sa,b5 as na,bH as ua,bI as ra,bJ as ia,a$ as ca,bL as da}from"./arco.d2aaf5b7.js";import"./chart.61872c57.js";import"./vue.ca65198a.js";const ma={class:"container"},pa={class:"action-icon"},_a={class:"action-icon"},fa={id:"tableSetting"},va={style:{"margin-right":"4px",cursor:"move"}},ba={class:"title"},ha={key:0,class:"circle"},ga={key:1,class:"circle pass"},ya={name:"UserList"},ka=Ee({...ya,setup($a){const K=H({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),X=async a=>{_(!0);try{await Ce(a),k()}catch{}finally{_(!1)}},E=()=>({user_id:d(),name:"",email:"",key:"",status:d(),created_at:[]}),{loading:W,setLoading:_}=Fe(!0),{t:n}=Ve(),L=d([]),s=d(E()),v=d([]),V=d([]),I=d("medium"),z={current:1,pageSize:10,showTotal:!0},D=H({...z}),Y=F(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),Z=F(()=>[{title:n("user.columns.userId"),dataIndex:"user_id",slotName:"user_id"},{title:n("user.columns.name"),dataIndex:"name",slotName:"name"},{title:n("user.columns.email"),dataIndex:"email",slotName:"email"},{title:n("user.columns.quota"),dataIndex:"quota",slotName:"quota"},{title:n("user.columns.status"),dataIndex:"status",slotName:"status"},{title:n("user.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at"},{title:n("user.columns.operations"),dataIndex:"operations",slotName:"operations"}]),ee=F(()=>[{label:n("user.dict.status.1"),value:1},{label:n("user.dict.status.2"),value:2}]),k=async(a={current:1,pageSize:10})=>{_(!0);try{const{data:l}=await Ie(a);L.value=l.items,D.current=a.current,D.total=l.paging.total}catch{}finally{_(!1)}},P=()=>{k({...z,...s.value})},ae=a=>{k({...z,...s.value,current:a})};k();const te=()=>{s.value=E()},le=(a,l)=>{I.value=a},oe=(a,l,u)=>{a?v.value.splice(u,0,l):v.value=V.value.filter(r=>r.dataIndex!==l.dataIndex)},x=(a,l,u,r=!1)=>{const m=r?T(a):a;return l>-1&&u>-1&&m.splice(l,1,m.splice(u,1,m[l]).pop()),m},se=a=>{a&&Ae(()=>{const l=document.getElementById("tableSetting");new Ne(l,{onEnd(u){const{oldIndex:r,newIndex:m}=u;x(v.value,r,m),x(V.value,r,m)}})})};Le(()=>Z.value,a=>{v.value=T(a),v.value.forEach((l,u)=>{l.checked=!0}),V.value=T(v.value)},{deep:!0,immediate:!0});const $=d(!1),A=d(),w=d({}),ne=async a=>{_(!0);try{w.value.user_id=a.user_id,$.value=!0}catch{}finally{_(!1)}},ue=async a=>{var u;if(await((u=A.value)==null?void 0:u.validate())){$.value=!0,a(!1);return}_(!0);try{await ze(w.value),Me.success(n("user.success.grantQuota")),a(),k()}catch{}finally{_(!1)}},re=()=>{$.value=!1};return(a,l)=>{const u=De,r=Oe,m=Qe,M=Re,b=Ge,c=He,U=je,ie=Je,ce=Ke,S=Xe,O=We,Q=Ye,de=Ze,g=ea,R=Ue,G=aa,me=ta,q=la,pe=Se,_e=oa,fe=sa,ve=qe,be=Be,he=na,ge=ua,ye=ra,ke=ia,$e=ca,we=da;return h(),y("div",ma,[e(m,{class:"container-breadcrumb"},{default:t(()=>[e(r,null,{default:t(()=>[e(u)]),_:1}),e(r,null,{default:t(()=>[p(i(a.$t("menu.user")),1)]),_:1}),e(r,null,{default:t(()=>[p(i(a.$t("menu.user.list")),1)]),_:1})]),_:1}),e(we,{class:"general-card",title:a.$t("menu.user.list"),bordered:!1},{extra:t(()=>[e($e,{visible:$.value,"onUpdate:visible":l[8]||(l[8]=o=>$.value=o),title:a.$t("user.form.title.grantQuota"),"ok-text":a.$t("user.button.save"),onCancel:re,onBeforeOk:ue},{default:t(()=>[e(O,{ref_key:"formRef",ref:A,model:w.value},{default:t(()=>[e(b,{field:"quota",label:a.$t("user.label.quota"),rules:[{required:!0,message:a.$t("user.error.quota.required")}]},{default:t(()=>[e(M,{modelValue:w.value.quota,"onUpdate:modelValue":l[7]||(l[7]=o=>w.value.quota=o),placeholder:a.$t("user.placeholder.quota")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),default:t(()=>[e(S,null,{default:t(()=>[e(c,{flex:1},{default:t(()=>[e(O,{model:s.value,"label-col-props":{span:6},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[e(S,{gutter:16},{default:t(()=>[e(c,{span:8},{default:t(()=>[e(b,{field:"user_id",label:a.$t("user.form.userId")},{default:t(()=>[e(M,{modelValue:s.value.user_id,"onUpdate:modelValue":l[0]||(l[0]=o=>s.value.user_id=o),placeholder:a.$t("user.form.userId.placeholder"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(c,{span:8},{default:t(()=>[e(b,{field:"name",label:a.$t("user.form.name")},{default:t(()=>[e(U,{modelValue:s.value.name,"onUpdate:modelValue":l[1]||(l[1]=o=>s.value.name=o),placeholder:a.$t("user.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(c,{span:8},{default:t(()=>[e(b,{field:"email",label:a.$t("user.form.email")},{default:t(()=>[e(U,{modelValue:s.value.email,"onUpdate:modelValue":l[2]||(l[2]=o=>s.value.email=o),placeholder:a.$t("user.form.email.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(c,{span:8},{default:t(()=>[e(b,{field:"key",label:a.$t("user.form.key")},{default:t(()=>[e(U,{modelValue:s.value.key,"onUpdate:modelValue":l[3]||(l[3]=o=>s.value.key=o),placeholder:a.$t("user.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(c,{span:8},{default:t(()=>[e(b,{field:"status",label:a.$t("user.form.status")},{default:t(()=>[e(ie,{modelValue:s.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>s.value.status=o),options:N(ee),placeholder:a.$t("user.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(c,{span:8},{default:t(()=>[e(b,{field:"created_at",label:a.$t("user.form.created_at")},{default:t(()=>[e(ce,{modelValue:s.value.created_at,"onUpdate:modelValue":l[5]||(l[5]=o=>s.value.created_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(Q,{style:{height:"84px"},direction:"vertical"}),e(c,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[e(G,{direction:"vertical",size:18},{default:t(()=>[e(g,{type:"primary",onClick:P},{icon:t(()=>[e(de)]),default:t(()=>[p(" "+i(a.$t("user.form.search")),1)]),_:1}),e(g,{onClick:te},{icon:t(()=>[e(R)]),default:t(()=>[p(" "+i(a.$t("user.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(Q,{style:{"margin-top":"0"}}),e(S,{style:{"margin-bottom":"16px"}},{default:t(()=>[e(c,{span:12},{default:t(()=>[e(G,null,{default:t(()=>[e(g,{type:"primary",onClick:l[6]||(l[6]=o=>a.$router.push({name:"UserCreate"}))},{icon:t(()=>[e(me)]),default:t(()=>[p(" "+i(a.$t("user.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(c,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[e(q,{content:a.$t("searchTable.actions.refresh")},{default:t(()=>[f("div",{class:"action-icon",onClick:P},[e(R,{size:"18"})])]),_:1},8,["content"]),e(fe,{onSelect:le},{content:t(()=>[(h(!0),y(j,null,J(N(Y),o=>(h(),Pe(_e,{key:o.value,value:o.value,class:xe({active:o.value===I.value})},{default:t(()=>[f("span",null,i(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[e(q,{content:a.$t("searchTable.actions.density")},{default:t(()=>[f("div",pa,[e(pe,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(q,{content:a.$t("searchTable.actions.columnSetting")},{default:t(()=>[e(ge,{trigger:"click",position:"bl",onPopupVisibleChange:se},{content:t(()=>[f("div",fa,[(h(!0),y(j,null,J(V.value,(o,C)=>(h(),y("div",{key:o.dataIndex,class:"setting"},[f("div",va,[e(be)]),f("div",null,[e(he,{modelValue:o.checked,"onUpdate:modelValue":B=>o.checked=B,onChange:B=>oe(B,o,C)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),f("div",ba,i(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:t(()=>[f("div",_a,[e(ve,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(ke,{"row-key":"id",loading:N(W),pagination:D,columns:v.value,data:L.value,bordered:!1,size:I.value,"row-selection":K,onPageChange:ae},{status:t(({record:o})=>[o.status===3?(h(),y("span",ha)):(h(),y("span",ga)),p(" "+i(a.$t(`user.dict.status.${o.status}`)),1)]),operations:t(({record:o})=>[e(g,{type:"text",size:"small",onClick:C=>ne({user_id:`${o.user_id}`})},{default:t(()=>[p(i(a.$t("user.columns.operations.grantQuota")),1)]),_:2},1032,["onClick"]),e(g,{type:"text",size:"small",onClick:C=>a.$router.push({name:"UserDetail",query:{id:`${o.id}`}})},{default:t(()=>[p(i(a.$t("user.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(ye,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:C=>X({id:`${o.id}`})},{default:t(()=>[e(g,{type:"text",size:"small"},{default:t(()=>[p(i(a.$t("user.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const Oa=Te(ka,[["__scopeId","data-v-9a8996c2"]]);export{Oa as default};
