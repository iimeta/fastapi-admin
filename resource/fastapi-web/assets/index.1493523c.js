import{u as He,k as Ge,p as je,y as Je,i as Ke,z as We,_ as Xe}from"./index.888b5da0.js";/* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css              */import{c as O,S as Ye}from"./sortable.esm.073617af.js";/* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as Ze,r as ne,e as p,c as A,w as ea,B as V,C as I,aH as a,aG as l,aL as u,aM as c,u as s,F as q,aJ as H,aI as G,aD as re,D as aa,bQ as r,n as ta,aW as ie,aK as la,aF as oa,aS as sa,b2 as ua,bC as na,b1 as ra,bD as ia,bB as da,bE as ca,b5 as ma,bF as pa,ab as _a,aU as fa,bi as va,bj as ba,bl as ga,bm as ha,b4 as ya,bG as $a,bT as qa,aT as wa,bH as ka,bI as Ca,bR as Va,bS as Da,a_ as xa,bA as Ua,bJ as za,g as Sa}from"./arco.7559a143.js";import{u as Ia}from"./loading.de2385ad.js";import{q as N}from"./common.4fed7ae4.js";import{s as Fa,q as Ba,a as Na,b as Ta,c as Qa,d as Pa}from"./admin_user.98b81c12.js";import{q as La}from"./model.8c4de245.js";import"./chart.1f478bf1.js";import"./vue.7fd3bfc3.js";import"./base.87fcf6e2.js";const Ra={class:"container"},Ea={class:"action-icon"},Ma={class:"action-icon"},Oa={id:"tableSetting"},Aa={style:{"margin-right":"4px",cursor:"move"}},Ha={class:"title"},Ga={name:"UserList"},ja=Ze({...Ga,setup(Ja){const{proxy:T}=Sa(),de=ne({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),j=p([]);(async()=>{try{const{data:e}=await La();j.value=e.items}catch{}})();const ce=async e=>{m(!0);try{await Fa(e),T.$message.success("\u5220\u9664\u6210\u529F"),U()}catch{}finally{m(!1)}},J=()=>({user_id:p(),name:"",email:"",quota_expires_at:[],status:p(),updated_at:[]}),{loading:me,setLoading:m}=Ia(!0),{t:i}=He(),f=p([]),d=p(J()),w=p([]),F=p([]),Q=p("medium"),D={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},B=ne({...D}),pe=A(()=>[{name:i("searchTable.size.mini"),value:"mini"},{name:i("searchTable.size.small"),value:"small"},{name:i("searchTable.size.medium"),value:"medium"},{name:i("searchTable.size.large"),value:"large"}]),_e=A(()=>[{title:i("user.columns.userId"),dataIndex:"user_id",slotName:"user_id",align:"center",width:80},{title:i("user.columns.name"),dataIndex:"name",slotName:"name",align:"center",ellipsis:!0,tooltip:!0},{title:i("user.columns.email"),dataIndex:"email",slotName:"email",align:"center",ellipsis:!0,tooltip:!0},{title:i("user.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center",ellipsis:!0,tooltip:!0},{title:i("user.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:i("user.columns.quota_expires_at"),dataIndex:"quota_expires_at",slotName:"quota_expires_at",align:"center",width:170},{title:i("user.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:i("user.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:i("user.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:318}]),fe=A(()=>[{label:i("user.dict.status.1"),value:1},{label:i("user.dict.status.2"),value:2}]),x=async(e={...D})=>{m(!0);try{const{data:o}=await Ba(e);f.value=o.items,B.current=e.current,B.pageSize=e.pageSize,B.total=o.paging.total}catch{}finally{m(!1)}},U=()=>{x({...D,...d.value})},ve=e=>{x({...D,...d.value,current:e})},be=e=>{D.pageSize=e,x({...D,...d.value})};x();const ge=()=>{d.value=J()},he=async e=>{m(!0);try{await Na(e),T.$message.success("\u64CD\u4F5C\u6210\u529F"),U()}catch{}finally{m(!1)}},ye=async e=>{m(!0);try{await Ta(e),T.$message.success("\u64CD\u4F5C\u6210\u529F"),U()}catch{}finally{m(!1)}},$e=(e,o)=>{Q.value=e},qe=(e,o,_)=>{e?w.value.splice(_,0,o):w.value=F.value.filter(b=>b.dataIndex!==o.dataIndex)},K=(e,o,_,b=!1)=>{const y=b?O(e):e;return o>-1&&_>-1&&y.splice(o,1,y.splice(_,1,y[o]).pop()),y},we=e=>{e&&ta(()=>{const o=document.getElementById("tableSetting");new Ye(o,{onEnd(_){const{oldIndex:b,newIndex:y}=_;K(w.value,b,y),K(F.value,b,y)}})})};ea(()=>_e.value,e=>{w.value=O(e),w.value.forEach((o,_)=>{o.checked=!0}),F.value=O(w.value)},{deep:!0,immediate:!0});const P=p(0),z=p(!1),S=p(!1),W=p(),n=p({}),X=p(),k=p({}),ke=async e=>{m(!0);try{P.value=0,n.value.quota=p(),n.value.user_id=e.user_id,n.value.quota_expires_at=e.quota_expires_at,z.value=!0}catch{}finally{m(!1)}},Ce=e=>{n.value.quota=e*5e5},Ve=async e=>{m(!0);try{k.value.user_id=e.user_id,e.models&&e.models.length>0&&e.models[0]!=="undefined"?k.value.models=e.models:k.value.models=[],S.value=!0}catch{}finally{m(!1)}},De=async e=>{var _;if(await((_=W.value)==null?void 0:_.validate())){z.value=!0,e(!1);return}m(!0);try{await Qa(n.value),ie.success(i("user.success.grantQuota")),e(),x()}catch{}finally{m(!1)}},xe=()=>{z.value=!1},Ue=async e=>{var _;if(await((_=X.value)==null?void 0:_.validate())){S.value=!0,e(!1);return}m(!0);try{await Pa(k.value),ie.success(i("user.success.models")),e(),x()}catch{}finally{m(!1)}},ze=()=>{S.value=!1};return(e,o)=>{const _=Ge,b=la,y=oa,Y=sa,g=ua,h=na,Z=ra,ee=ia,ae=da,L=ca,R=ma,te=pa,Se=_a,$=fa,le=je,oe=va,E=ba,Ie=Je,Fe=ga,Be=ha,Ne=Ke,Te=We,Qe=ya,Pe=$a,se=qa,Le=wa,Re=ka,Ee=Ca,C=Va,Me=Da,ue=xa,Oe=Ua,Ae=za;return V(),I("div",Ra,[a(y,{class:"container-breadcrumb"},{default:l(()=>[a(b,null,{default:l(()=>[a(_)]),_:1}),a(b,null,{default:l(()=>[u(c(e.$t("menu.user")),1)]),_:1}),a(b,null,{default:l(()=>[u(c(e.$t("menu.user.list")),1)]),_:1})]),_:1}),a(Ae,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:l(()=>[a(L,null,{default:l(()=>[a(h,{flex:1},{default:l(()=>[a(R,{model:d.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:l(()=>[a(L,{gutter:16},{default:l(()=>[a(h,{span:8},{default:l(()=>[a(g,{field:"user_id",label:e.$t("user.form.userId")},{default:l(()=>[a(Y,{modelValue:d.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=t=>d.value.user_id=t),placeholder:e.$t("user.form.userId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(h,{span:8},{default:l(()=>[a(g,{field:"name",label:e.$t("user.form.name")},{default:l(()=>[a(Z,{modelValue:d.value.name,"onUpdate:modelValue":o[1]||(o[1]=t=>d.value.name=t),placeholder:e.$t("user.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(h,{span:8},{default:l(()=>[a(g,{field:"email",label:e.$t("user.form.email")},{default:l(()=>[a(Z,{modelValue:d.value.email,"onUpdate:modelValue":o[2]||(o[2]=t=>d.value.email=t),placeholder:e.$t("user.form.email.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(h,{span:8},{default:l(()=>[a(g,{field:"quota_expires_at",label:e.$t("user.form.quota_expires_at")},{default:l(()=>[a(ee,{modelValue:d.value.quota_expires_at,"onUpdate:modelValue":o[3]||(o[3]=t=>d.value.quota_expires_at=t),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1}),a(h,{span:8},{default:l(()=>[a(g,{field:"status",label:e.$t("user.form.status")},{default:l(()=>[a(ae,{modelValue:d.value.status,"onUpdate:modelValue":o[4]||(o[4]=t=>d.value.status=t),options:s(fe),placeholder:e.$t("user.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(h,{span:8},{default:l(()=>[a(g,{field:"updated_at",label:e.$t("user.form.updated_at")},{default:l(()=>[a(ee,{modelValue:d.value.updated_at,"onUpdate:modelValue":o[5]||(o[5]=t=>d.value.updated_at=t),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(te,{style:{height:"84px"},direction:"vertical"}),a(h,{flex:"86px",style:{"text-align":"right"}},{default:l(()=>[a(oe,{direction:"vertical",size:18},{default:l(()=>[a($,{type:"primary",onClick:U},{icon:l(()=>[a(Se)]),default:l(()=>[u(" "+c(e.$t("user.form.search")),1)]),_:1}),a($,{onClick:ge},{icon:l(()=>[a(le)]),default:l(()=>[u(" "+c(e.$t("user.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(te,{style:{"margin-top":"0","margin-bottom":"16px"}}),a(L,{style:{"margin-bottom":"16px"}},{default:l(()=>[a(h,{span:12},{default:l(()=>[a(oe,null,{default:l(()=>[a($,{type:"primary",onClick:o[6]||(o[6]=t=>e.$router.push({name:"UserCreate"}))},{default:l(()=>[u(c(e.$t("user.operation.create")),1)]),_:1})]),_:1})]),_:1}),a(h,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:l(()=>[a(E,{content:e.$t("searchTable.actions.refresh")},{default:l(()=>[q("div",{class:"action-icon",onClick:U},[a(le,{size:"18"})])]),_:1},8,["content"]),a(Be,{onSelect:$e},{content:l(()=>[(V(!0),I(H,null,G(s(pe),t=>(V(),re(Fe,{key:t.value,value:t.value,class:aa({active:t.value===Q.value})},{default:l(()=>[q("span",null,c(t.name),1)]),_:2},1032,["value","class"]))),128))]),default:l(()=>[a(E,{content:e.$t("searchTable.actions.density")},{default:l(()=>[q("div",Ea,[a(Ie,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(E,{content:e.$t("searchTable.actions.columnSetting")},{default:l(()=>[a(Pe,{trigger:"click",position:"bl",onPopupVisibleChange:we},{content:l(()=>[q("div",Oa,[(V(!0),I(H,null,G(F.value,(t,v)=>(V(),I("div",{key:t.dataIndex,class:"setting"},[q("div",Aa,[a(Te)]),q("div",null,[a(Qe,{modelValue:t.checked,"onUpdate:modelValue":M=>t.checked=M,onChange:M=>qe(M,t,v)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),q("div",Ha,c(t.title==="#"?"\u5E8F\u5217\u53F7":t.title),1)]))),128))])]),default:l(()=>[q("div",Ma,[a(Ne,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(Ee,{"row-key":"id",loading:s(me),pagination:B,columns:w.value,data:f.value,bordered:!1,size:Q.value,"row-selection":de,onPageChange:ve,onPageSizeChange:be},{quota:l(({record:t})=>[u(c(t.quota>0?`$${s(N)(t.quota)}`:t.quota<0?`-$${s(N)(-t.quota)}`:"$0.00"),1)]),used_quota:l(({record:t})=>[u(" $"+c(t.used_quota>0?s(N)(t.used_quota):"0.00"),1)]),quota_expires_at:l(({rowIndex:t})=>[a(se,{modelValue:f.value[t].quota_expires_at,"onUpdate:modelValue":v=>f.value[t].quota_expires_at=v,placeholder:e.$t("user.columns.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":v=>s(r)(v).isBefore(s(r)()),"show-time":"",shortcuts:[{label:"1",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(1,"day")},{label:"7",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(7,"day")},{label:"15",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(15,"day")},{label:"30",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(30,"day")},{label:"90",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(90,"day")},{label:"180",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(180,"day")},{label:"365",value:()=>s(r)(new Date(f.value[t].quota_expires_at||new Date)).add(365,"day")}],onChange:v=>he({id:`${f.value[t].id}`,quota_expires_at:`${f.value[t].quota_expires_at}`})},{default:l(()=>[a($,{style:{width:"150px"}},{default:l(()=>[u(c(f.value[t].quota_expires_at||e.$t("user.columns.placeholder.quota_expires_at")),1)]),_:2},1024)]),_:2},1032,["modelValue","onUpdate:modelValue","placeholder","disabled-date","shortcuts","onChange"])]),remark:l(({record:t})=>[u(c(t.remark||"-"),1)]),status:l(({record:t})=>[a(Le,{modelValue:t.status,"onUpdate:modelValue":v=>t.status=v,"checked-value":1,"unchecked-value":2,onChange:v=>ye({id:`${t.id}`,status:Number(`${t.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:l(({record:t})=>[a($,{type:"text",size:"small",onClick:v=>ke({user_id:`${t.user_id}`,quota_expires_at:`${t.quota_expires_at}`})},{default:l(()=>[u(c(e.$t("user.columns.operations.grantQuota")),1)]),_:2},1032,["onClick"]),a($,{type:"text",size:"small",onClick:v=>Ve({user_id:`${t.user_id}`,models:`${t.models}`.split(",")})},{default:l(()=>[u(c(e.$t("user.columns.operations.models")),1)]),_:2},1032,["onClick"]),a($,{type:"text",size:"small",onClick:v=>e.$router.push({name:"UserDetail",query:{id:`${t.id}`}})},{default:l(()=>[u(c(e.$t("user.columns.operations.view")),1)]),_:2},1032,["onClick"]),a($,{type:"text",size:"small",onClick:v=>e.$router.push({name:"UserUpdate",query:{id:`${t.id}`}})},{default:l(()=>[u(c(e.$t("user.columns.operations.update")),1)]),_:2},1032,["onClick"]),a(Re,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:v=>ce({id:`${t.id}`})},{default:l(()=>[a($,{type:"text",size:"small"},{default:l(()=>[u(c(e.$t("user.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(ue,{visible:z.value,"onUpdate:visible":o[10]||(o[10]=t=>z.value=t),title:e.$t("user.form.title.grantQuota"),"ok-text":e.$t("user.button.save"),onCancel:xe,onBeforeOk:De},{default:l(()=>[a(R,{ref_key:"formRef",ref:W,model:n.value},{default:l(()=>[a(g,{field:"quota",label:e.$t("user.label.quota"),rules:[{required:!0,message:e.$t("user.error.quota.required")}]},{default:l(()=>[a(Y,{modelValue:n.value.quota,"onUpdate:modelValue":o[7]||(o[7]=t=>n.value.quota=t),placeholder:e.$t("user.placeholder.grant_quota"),precision:0,min:-9999999999999,max:9999999999999,style:{"margin-right":"10px"}},null,8,["modelValue","placeholder"]),q("div",null," $"+c(n.value.quota?s(N)(n.value.quota):"0.00"),1)]),_:1},8,["label","rules"]),a(g,null,{default:l(()=>[a(Me,{modelValue:P.value,"onUpdate:modelValue":o[8]||(o[8]=t=>P.value=t),type:"button",onChange:Ce},{default:l(()=>[a(C,{value:1},{default:l(()=>[u(" $1 ")]),_:1}),a(C,{value:5},{default:l(()=>[u(" $5 ")]),_:1}),a(C,{value:10},{default:l(()=>[u(" $10 ")]),_:1}),a(C,{value:20},{default:l(()=>[u(" $20 ")]),_:1}),a(C,{value:100},{default:l(()=>[u(" $100 ")]),_:1}),a(C,{value:500},{default:l(()=>[u(" $500 ")]),_:1}),a(C,{value:1e3},{default:l(()=>[u(" $1000 ")]),_:1})]),_:1},8,["modelValue","onChange"])]),_:1}),a(g,{field:"quota_expires_at",label:e.$t("user.label.quota_expires_at")},{default:l(()=>[a(se,{modelValue:n.value.quota_expires_at,"onUpdate:modelValue":o[9]||(o[9]=t=>n.value.quota_expires_at=t),placeholder:e.$t("user.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":t=>s(r)(t).isBefore(s(r)()),position:"tl",style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(1,"day")},{label:"7",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(7,"day")},{label:"15",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(15,"day")},{label:"30",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(30,"day")},{label:"90",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(90,"day")},{label:"180",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(180,"day")},{label:"365",value:()=>s(r)(new Date(n.value.quota_expires_at||new Date)).add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"]),a(ue,{visible:S.value,"onUpdate:visible":o[12]||(o[12]=t=>S.value=t),title:e.$t("user.form.title.models"),"ok-text":e.$t("user.button.save"),width:700,onCancel:ze,onBeforeOk:Ue},{default:l(()=>[a(R,{ref_key:"modelsFormRef",ref:X,model:k.value},{default:l(()=>[a(g,{field:"models",label:e.$t("user.label.models")},{default:l(()=>[a(ae,{modelValue:k.value.models,"onUpdate:modelValue":o[11]||(o[11]=t=>k.value.models=t),placeholder:e.$t("user.placeholder.models"),multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(V(!0),I(H,null,G(j.value,t=>(V(),re(Oe,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),_:1})])}}});const wt=Xe(ja,[["__scopeId","data-v-826fd51a"]]);export{wt as default};
