import{u as Fe,j as Le,o as Pe,x as xe,I as Me,y as Oe,_ as Ee}from"./index.525e8576.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                */import{c as x,S as Ae}from"./sortable.esm.2109e0e3.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as Qe,r as Z,e as d,c as M,w as Re,B as f,C as $,aH as a,aG as t,aL as m,aM as i,aJ as O,aI as E,aD as ee,u as A,F as g,D as He,n as je,aW as ae,aK as Ge,aF as Je,aS as Ke,b2 as We,bA as Xe,b1 as Ye,bC as Ze,bD as ea,bE as aa,b5 as ta,bF as la,ab as oa,aU as sa,bi as na,a5 as ua,bj as ra,bl as ia,bm as da,b4 as ca,bG as ma,bH as pa,bI as _a,a_ as fa,bB as va,bJ as ba}from"./arco.aed15247.js";import{u as ga}from"./loading.b5911e1d.js";import{s as ha,q as ya,a as $a,b as ka,c as wa}from"./admin_user.56f10fe7.js";import{q as Ca}from"./model.570552fc.js";import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const Va={class:"container"},Ia={class:"action-icon"},za={class:"action-icon"},Sa={id:"tableSetting"},Ua={style:{"margin-right":"4px",cursor:"move"}},qa={class:"title"},Da={key:0,class:"circle red"},Ba={key:1,class:"circle"},Na={name:"UserList"},Ta=Qe({...Na,setup(Fa){const te=Z({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),Q=d([]);(async()=>{try{const{data:e}=await Ca();Q.value=e.items}catch{}})();const le=async e=>{r(!0);try{await ha(e),q()}catch{}finally{r(!1)}},R=()=>({user_id:d(),name:"",email:"",key:"",status:d(),created_at:[]}),{loading:oe,setLoading:r}=ga(!0),{t:n}=Fe(),H=d([]),s=d(R()),h=d([]),S=d([]),B=d("medium"),k={current:1,pageSize:10,showTotal:!0,showPageSize:!0},U=Z({...k}),se=M(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),ne=M(()=>[{title:n("user.columns.userId"),dataIndex:"user_id",slotName:"user_id",align:"center"},{title:n("user.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:n("user.columns.email"),dataIndex:"email",slotName:"email",align:"center"},{title:n("user.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center"},{title:n("user.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:80},{title:n("user.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:170},{title:n("user.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:320}]),ue=M(()=>[{label:n("user.dict.status.1"),value:1},{label:n("user.dict.status.2"),value:2}]),w=async(e={...k})=>{r(!0);try{const{data:o}=await ya(e);H.value=o.items,U.current=e.current,U.pageSize=e.pageSize,U.total=o.paging.total}catch{}finally{r(!1)}},q=()=>{w({...k,...s.value})},re=e=>{w({...k,...s.value,current:e})},ie=e=>{k.pageSize=e,w({...k,...s.value})};w();const de=()=>{s.value=R()},ce=async e=>{r(!0);try{e.status=e.status===1?2:1,await $a(e),q()}catch{}finally{r(!1)}},me=(e,o)=>{B.value=e},pe=(e,o,u)=>{e?h.value.splice(u,0,o):h.value=S.value.filter(c=>c.dataIndex!==o.dataIndex)},j=(e,o,u,c=!1)=>{const _=c?x(e):e;return o>-1&&u>-1&&_.splice(o,1,_.splice(u,1,_[o]).pop()),_},_e=e=>{e&&je(()=>{const o=document.getElementById("tableSetting");new Ae(o,{onEnd(u){const{oldIndex:c,newIndex:_}=u;j(h.value,c,_),j(S.value,c,_)}})})};Re(()=>ne.value,e=>{h.value=x(e),h.value.forEach((o,u)=>{o.checked=!0}),S.value=x(h.value)},{deep:!0,immediate:!0});const V=d(!1),I=d(!1),D=d(),z=d({}),y=d({}),fe=async e=>{r(!0);try{z.value.user_id=e.user_id,V.value=!0}catch{}finally{r(!1)}},ve=async e=>{r(!0);try{y.value.user_id=e.user_id,e.models&&e.models.length>0&&e.models[0]!=="undefined"?y.value.models=e.models:y.value.models=[],I.value=!0}catch{}finally{r(!1)}},be=async e=>{var u;if(await((u=D.value)==null?void 0:u.validate())){V.value=!0,e(!1);return}r(!0);try{await ka(z.value),ae.success(n("user.success.grantQuota")),e(),w()}catch{}finally{r(!1)}},ge=()=>{V.value=!1},he=async e=>{var u;if(await((u=D.value)==null?void 0:u.validate())){I.value=!0,e(!1);return}r(!0);try{await wa(y.value),ae.success(n("user.success.models")),e(),w()}catch{}finally{r(!1)}},ye=()=>{I.value=!1};return(e,o)=>{const u=Le,c=Ge,_=Je,G=Ke,v=We,p=Xe,N=Ye,J=Ze,$e=ea,T=aa,F=ta,K=la,ke=oa,b=sa,W=Pe,X=na,we=ua,L=ra,Ce=xe,Ve=ia,Ie=da,ze=Me,Se=Oe,Ue=ca,qe=ma,De=pa,Be=_a,Y=fa,Ne=va,Te=ba;return f(),$("div",Va,[a(_,{class:"container-breadcrumb"},{default:t(()=>[a(c,null,{default:t(()=>[a(u)]),_:1}),a(c,null,{default:t(()=>[m(i(e.$t("menu.user")),1)]),_:1}),a(c,null,{default:t(()=>[m(i(e.$t("menu.user.list")),1)]),_:1})]),_:1}),a(Te,{class:"general-card",title:e.$t("menu.user.list"),bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"0 20px 20px"}},{extra:t(()=>[a(Y,{visible:V.value,"onUpdate:visible":o[8]||(o[8]=l=>V.value=l),title:e.$t("user.form.title.grantQuota"),"ok-text":e.$t("user.button.save"),onCancel:ge,onBeforeOk:be},{default:t(()=>[a(F,{ref_key:"formRef",ref:D,model:z.value},{default:t(()=>[a(v,{field:"quota",label:e.$t("user.label.quota"),rules:[{required:!0,message:e.$t("user.error.quota.required")}]},{default:t(()=>[a(G,{modelValue:z.value.quota,"onUpdate:modelValue":o[7]||(o[7]=l=>z.value.quota=l),placeholder:e.$t("user.placeholder.quota"),precision:0,min:-9999999999999,max:9999999999999},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"]),a(Y,{visible:I.value,"onUpdate:visible":o[10]||(o[10]=l=>I.value=l),title:e.$t("user.form.title.models"),"ok-text":e.$t("user.button.save"),onCancel:ye,onBeforeOk:he},{default:t(()=>[a(F,{ref_key:"formRef",ref:D,model:y.value},{default:t(()=>[a(v,{field:"models",label:e.$t("user.label.models")},{default:t(()=>[a(J,{modelValue:y.value.models,"onUpdate:modelValue":o[9]||(o[9]=l=>y.value.models=l),placeholder:e.$t("user.placeholder.models"),"max-tag-count":15,multiple:"","allow-clear":""},{default:t(()=>[(f(!0),$(O,null,E(Q.value,l=>(f(),ee(Ne,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),default:t(()=>[a(T,null,{default:t(()=>[a(p,{flex:1},{default:t(()=>[a(F,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[a(T,{gutter:16},{default:t(()=>[a(p,{span:8},{default:t(()=>[a(v,{field:"user_id",label:e.$t("user.form.userId")},{default:t(()=>[a(G,{modelValue:s.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=l=>s.value.user_id=l),placeholder:e.$t("user.form.userId.placeholder"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(p,{span:8},{default:t(()=>[a(v,{field:"name",label:e.$t("user.form.name")},{default:t(()=>[a(N,{modelValue:s.value.name,"onUpdate:modelValue":o[1]||(o[1]=l=>s.value.name=l),placeholder:e.$t("user.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(p,{span:8},{default:t(()=>[a(v,{field:"email",label:e.$t("user.form.email")},{default:t(()=>[a(N,{modelValue:s.value.email,"onUpdate:modelValue":o[2]||(o[2]=l=>s.value.email=l),placeholder:e.$t("user.form.email.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(p,{span:8},{default:t(()=>[a(v,{field:"key",label:e.$t("user.form.key")},{default:t(()=>[a(N,{modelValue:s.value.key,"onUpdate:modelValue":o[3]||(o[3]=l=>s.value.key=l),placeholder:e.$t("user.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(p,{span:8},{default:t(()=>[a(v,{field:"status",label:e.$t("user.form.status")},{default:t(()=>[a(J,{modelValue:s.value.status,"onUpdate:modelValue":o[4]||(o[4]=l=>s.value.status=l),options:A(ue),placeholder:e.$t("user.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(p,{span:8},{default:t(()=>[a(v,{field:"created_at",label:e.$t("user.form.created_at")},{default:t(()=>[a($e,{modelValue:s.value.created_at,"onUpdate:modelValue":o[5]||(o[5]=l=>s.value.created_at=l),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(K,{style:{height:"84px"},direction:"vertical"}),a(p,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[a(X,{direction:"vertical",size:18},{default:t(()=>[a(b,{type:"primary",onClick:q},{icon:t(()=>[a(ke)]),default:t(()=>[m(" "+i(e.$t("user.form.search")),1)]),_:1}),a(b,{onClick:de},{icon:t(()=>[a(W)]),default:t(()=>[m(" "+i(e.$t("user.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(K,{style:{"margin-top":"0"}}),a(T,{style:{"margin-bottom":"16px"}},{default:t(()=>[a(p,{span:12},{default:t(()=>[a(X,null,{default:t(()=>[a(b,{type:"primary",onClick:o[6]||(o[6]=l=>e.$router.push({name:"UserCreate"}))},{icon:t(()=>[a(we)]),default:t(()=>[m(" "+i(e.$t("user.operation.create")),1)]),_:1})]),_:1})]),_:1}),a(p,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:t(()=>[a(L,{content:e.$t("searchTable.actions.refresh")},{default:t(()=>[g("div",{class:"action-icon",onClick:q},[a(W,{size:"18"})])]),_:1},8,["content"]),a(Ie,{onSelect:me},{content:t(()=>[(f(!0),$(O,null,E(A(se),l=>(f(),ee(Ve,{key:l.value,value:l.value,class:He({active:l.value===B.value})},{default:t(()=>[g("span",null,i(l.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[a(L,{content:e.$t("searchTable.actions.density")},{default:t(()=>[g("div",Ia,[a(Ce,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(L,{content:e.$t("searchTable.actions.columnSetting")},{default:t(()=>[a(qe,{trigger:"click",position:"bl",onPopupVisibleChange:_e},{content:t(()=>[g("div",Sa,[(f(!0),$(O,null,E(S.value,(l,C)=>(f(),$("div",{key:l.dataIndex,class:"setting"},[g("div",Ua,[a(Se)]),g("div",null,[a(Ue,{modelValue:l.checked,"onUpdate:modelValue":P=>l.checked=P,onChange:P=>pe(P,l,C)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),g("div",qa,i(l.title==="#"?"\u5E8F\u5217\u53F7":l.title),1)]))),128))])]),default:t(()=>[g("div",za,[a(ze,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(Be,{"row-key":"id",loading:A(oe),pagination:U,columns:h.value,data:H.value,bordered:!1,size:B.value,"row-selection":te,onPageChange:re,onPageSizeChange:ie},{status:t(({record:l})=>[l.status===2?(f(),$("span",Da)):(f(),$("span",Ba)),m(" "+i(e.$t(`user.dict.status.${l.status}`)),1)]),quota:t(({record:l})=>[m(i(l.quota.toLocaleString()),1)]),operations:t(({record:l})=>[a(b,{type:"text",size:"small",onClick:C=>fe({user_id:`${l.user_id}`})},{default:t(()=>[m(i(e.$t("user.columns.operations.grantQuota")),1)]),_:2},1032,["onClick"]),a(b,{type:"text",size:"small",onClick:C=>ve({user_id:`${l.user_id}`,models:`${l.models}`.split(",")})},{default:t(()=>[m(i(e.$t("user.columns.operations.models")),1)]),_:2},1032,["onClick"]),a(b,{type:"text",size:"small",onClick:C=>e.$router.push({name:"UserDetail",query:{id:`${l.id}`}})},{default:t(()=>[m(i(e.$t("user.columns.operations.view")),1)]),_:2},1032,["onClick"]),a(b,{type:"text",size:"small",onClick:C=>ce({id:`${l.id}`,status:Number(`${l.status}`)})},{default:t(()=>[m(i(e.$t(`user.columns.operations.status.${l.status}`)),1)]),_:2},1032,["onClick"]),a(De,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:C=>le({id:`${l.id}`})},{default:t(()=>[a(b,{type:"text",size:"small"},{default:t(()=>[m(i(e.$t("user.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const nt=Ee(Ta,[["__scopeId","data-v-0fe45f1b"]]);export{nt as default};
