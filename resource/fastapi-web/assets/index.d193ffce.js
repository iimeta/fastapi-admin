import{_ as pe,u as Ge,G as Je,p as Oe,y as Ke,i as Qe,z as We}from"./index.853b9bb5.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as fe,e as $,w as ve,bS as ye,B as t,C as d,bu as h,aD as n,aG as e,aH as a,u as r,aL as _,aM as s,aE as ke,g as Xe,bJ as Ze,bK as ea,ad as aa,bL as ta,bM as he,bN as la,bi as be,bb as oa,b6 as na,r as me,c as te,aJ as N,aI as M,F as S,D as sa,bT as ce,n as ua,aK as ia,aF as da,bA as ra,bB as _a,b2 as ma,bC as ca,b1 as pa,bU as fa,bD as va,aS as ya,b5 as ka,bE as ha,ab as ba,aU as ga,bj as wa,bl as $a,bm as Ca,b4 as Va,bF as Da,c2 as Fa,bH as Ia,aV as Ba,bI as Ea}from"./arco.a9260898.js";import{u as ge}from"./loading.1f346a94.js";import{q as xa}from"./common.df364eef.js";import{V as Aa,d as Sa,i as qa,e as za}from"./styles.b61f5bc2.js";import{g as La}from"./app.394bbdd2.js";import{c as le,S as Ua}from"./sortable.esm.a0dfbf42.js";import{q as Na}from"./model.2f2e871f.js";import{f as Ma}from"./agent.4e32dc47.js";/* empty css               *//* empty css                *//* empty css                */import{u as Pa}from"./vue.ad52ddbe.js";import"./chart.d103b168.js";const Ta={style:{margin:"10px 0 30px 10px"}},Ya={key:1},Ha={key:1},Ra={key:1},ja={key:1},Ga={key:1},Ja={key:1},Oa={key:1},Ka={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},Qa={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},Wa={key:1},Xa={key:1},Za={key:1},et={key:1},at={key:1},tt={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},lt={key:1},ot={key:1},nt={key:1},st={key:1},ut={key:1},it={key:1},dt={key:1},rt={key:1},_t={key:1},mt={key:1},ct={key:1},pt={key:1},ft={key:1},vt={key:1},yt={key:1},kt={key:1},ht={key:1},bt={key:1},gt={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},wt={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},$t={key:1},Ct={key:1},Vt={key:1},Dt={key:1},Ft={key:1},It={key:1},Bt={key:1},Et={key:1},xt={key:1},At={key:1},St={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},qt={key:1},zt={name:"ImageDetail"},Lt=fe({...zt,props:{id:{type:String,default:""}},setup(oe){const E=oe,{loading:p,setLoading:P}=ge(!0),o=$({}),{copy:z,copied:Q}=Pa(),{proxy:R}=Xe();(async(k={id:E.id})=>{P(!0);try{const{data:v}=await Sa(k);o.value=v}catch{}finally{P(!1)}})();const T=k=>{z(k)};ve(Q,()=>{Q.value&&R.$message.success("\u590D\u5236\u6210\u529F")});const Y=async(k,v)=>{const{data:m}=await qa({id:k,field:v});z(m.value)};return(k,v)=>{const m=Ze,u=ea,D=aa,c=ta,y=he,I=la,L=be,W=oa,j=na,q=ye("permission");return t(),d("div",Ta,[h((t(),n(I,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:e(()=>[a(c,{label:"Trace ID",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ya,[_(s(o.value.trace_id)+" ",1),a(D,{class:"copy-btn",onClick:v[0]||(v[0]=g=>T(o.value.trace_id))})]))]),_:1}),a(c,{label:"Host",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ha,s(o.value.host||"-"),1))]),_:1}),a(c,{label:"\u8C03\u7528\u5BC6\u94A5",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ra,[_(s(o.value.creator)+" ",1),a(D,{class:"copy-btn",onClick:v[1]||(v[1]=g=>Y(o.value.id,"creator"))})]))]),_:1}),a(c,{label:"\u7528\u6237ID"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",ja,s(o.value.user_id||"-"),1))]),_:1}),a(c,{label:"\u5E94\u7528ID",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",Ga,s(o.value.app_id||"-"),1))]),_:1}),a(c,{label:"\u6A21\u578B"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ja,s(o.value.model||"-"),1))]),_:1}),a(c,{label:"\u6A21\u578B\u7C7B\u578B"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Oa,s(k.$t(`dict.model_type.${o.value.type}`)),1))]),_:1}),a(c,{label:"\u63D0\u793A\u8BCD",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ka,s(o.value.prompt||"-"),1))]),_:1}),a(c,{label:"\u56DE\u7B54"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Qa,s(o.value.image_data||"-"),1))]),_:1}),a(c,{label:"\u82B1\u8D39\u4EE4\u724C\u6570",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Wa,s(o.value.total_tokens?o.value.total_tokens:o.value.status===1&&o.value.text_quota.billing_method===2?0:"-"),1))]),_:1}),a(c,{label:"\u603B\u8017\u65F6",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Xa,[o.value.total_time>18e4?(t(),n(y,{key:0,color:"red"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):o.value.total_time>12e4?(t(),n(y,{key:1,color:"orange"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):o.value.total_time>9e4?(t(),n(y,{key:2,color:"gold"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):(t(),n(y,{key:3,color:"green"},{default:e(()=>[_(s(o.value.total_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(c,{label:"\u7ED3\u679C",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Za,[o.value.status===1?(t(),n(y,{key:0,color:"green"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):o.value.status===2?(t(),n(y,{key:1,color:"gold"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):o.value.status===3?(t(),n(y,{key:2,color:"orange"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):(t(),n(y,{key:3,color:"red"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1}),a(c,{label:"\u5BA2\u6237\u7AEFIP",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",et,s(o.value.client_ip||"-"),1))]),_:1}),a(c,{label:"\u8BF7\u6C42\u65F6\u95F4",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",at,s(o.value.req_time||"-"),1))]),_:1}),a(c,{label:"\u9519\u8BEF\u4FE1\u606F",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",tt,s(o.value.err_msg||"-"),1))]),_:1})]),_:1})),[[q,["user"]]]),h((t(),n(I,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:e(()=>[a(c,{label:"Trace ID",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",lt,[_(s(o.value.trace_id)+" ",1),a(D,{class:"copy-btn",onClick:v[2]||(v[2]=g=>T(o.value.trace_id))})]))]),_:1}),a(c,{label:"Host"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",ot,s(o.value.host||"-"),1))]),_:1}),a(c,{label:"\u8C03\u7528\u5BC6\u94A5",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",nt,[_(s(o.value.is_smart_match?"-":o.value.creator)+" ",1),o.value.is_smart_match?ke("",!0):(t(),n(D,{key:0,class:"copy-btn",onClick:v[3]||(v[3]=g=>Y(o.value.id,"creator"))}))]))]),_:1}),a(c,{label:"\u7528\u6237ID"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",st,s(o.value.is_smart_match?"-":o.value.user_id||"-"),1))]),_:1}),a(c,{label:"\u5E94\u7528ID",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",ut,s(o.value.is_smart_match?"-":o.value.app_id||"-"),1))]),_:1}),a(c,{label:"\u516C\u53F8"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",it,s(o.value.corp_name),1))]),_:1}),a(c,{label:"\u6A21\u578B\u7C7B\u578B"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",dt,s(k.$t(`dict.model_type.${o.value.type}`)),1))]),_:1}),a(c,{label:"\u8BF7\u6C42\u6A21\u578B\u540D\u79F0"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",rt,s(o.value.name||"-"),1))]),_:1}),a(c,{label:"\u8BF7\u6C42\u6A21\u578B"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",_t,s(o.value.model||"-"),1))]),_:1}),a(c,{label:"\u771F\u5B9E\u6A21\u578B\u540D\u79F0"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",mt,s(o.value.real_model_name),1))]),_:1}),a(c,{label:"\u771F\u5B9E\u6A21\u578B"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",ct,s(o.value.real_model),1))]),_:1}),a(c,{label:"\u540E\u5907\u4EE3\u7406"},{default:e(()=>{var g,F;return[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",pt,s(((F=(g=o.value)==null?void 0:g.fallback_config)==null?void 0:F.model_agent_name)||"-"),1))]}),_:1}),a(c,{label:"\u540E\u5907\u6A21\u578B"},{default:e(()=>{var g,F;return[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",ft,s(((F=(g=o.value)==null?void 0:g.fallback_config)==null?void 0:F.model)||"-"),1))]}),_:1}),a(c,{label:"\u542F\u7528\u8F6C\u53D1"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",vt,s(k.$t(`chat.dict.is_enable_forward.${o.value.is_enable_forward||!1}`)),1))]),_:1}),a(c,{label:"\u8F6C\u53D1\u89C4\u5219"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",yt,s(o.value.is_enable_forward?k.$t(`chat.dict.forward_rule.${o.value.forward_config.forward_rule||"1"}`):"-"),1))]),_:1}),a(c,{label:"\u542F\u7528\u4EE3\u7406",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",kt,s(k.$t(`chat.dict.is_enable_model_agent.${o.value.is_enable_model_agent||!1}`)),1))]),_:1}),a(c,{label:"\u4EE3\u7406\u540D\u79F0"},{default:e(()=>{var g,F;return[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",ht,s(((F=(g=o.value)==null?void 0:g.model_agent)==null?void 0:F.name)||"-"),1))]}),_:1}),a(c,{label:"\u5BC6\u94A5"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",bt,[_(s(o.value.key?o.value.key.length>0?o.value.key.substr(0,o.value.key.length/2>10?10:o.value.key.length/2)+"************************************"+o.value.key.substr(-(o.value.key.length/2>5?5:o.value.key.length/2)):o.value.key:"-")+" ",1),a(D,{class:"copy-btn",onClick:v[4]||(v[4]=g=>Y(o.value.id,"key"))})]))]),_:1}),a(c,{label:"\u63D0\u793A\u8BCD",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",gt,s(o.value.prompt||"-"),1))]),_:1}),a(c,{label:"\u56DE\u7B54"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",wt,s(o.value.image_data||"-"),1))]),_:1}),a(c,{label:"\u8BA1\u8D39\u65B9\u5F0F",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",$t,s(k.$t("chat.dict.billing_method.2")),1))]),_:1}),a(c,{label:"\u82B1\u8D39\u4EE4\u724C\u6570"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ct,s(o.value.total_tokens?o.value.total_tokens:o.value.status===1&&o.value.text_quota.billing_method===2?0:"-"),1))]),_:1}),a(c,{label:"\u603B\u8017\u65F6",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Vt,[o.value.total_time>12e4?(t(),n(y,{key:0,color:"red"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):o.value.total_time>9e4?(t(),n(y,{key:1,color:"orange"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):o.value.total_time>6e4?(t(),n(y,{key:2,color:"gold"},{default:e(()=>[_(s(o.value.total_time)+" ms ",1)]),_:1})):(t(),n(y,{key:3,color:"green"},{default:e(()=>[_(s(o.value.total_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(c,{label:"\u5185\u8017"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Dt,[o.value.internal_time>500?(t(),n(y,{key:0,color:"red"},{default:e(()=>[_(s(o.value.internal_time)+" ms ",1)]),_:1})):o.value.internal_time>300?(t(),n(y,{key:1,color:"orange"},{default:e(()=>[_(s(o.value.internal_time)+" ms ",1)]),_:1})):o.value.internal_time>100?(t(),n(y,{key:2,color:"gold"},{default:e(()=>[_(s(o.value.internal_time)+" ms ",1)]),_:1})):(t(),n(y,{key:3,color:"green"},{default:e(()=>[_(s(o.value.internal_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(c,{label:"\u7ED3\u679C",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",Ft,[o.value.status===1?(t(),n(y,{key:0,color:"green"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):o.value.status===2?(t(),n(y,{key:1,color:"gold"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):o.value.status===3?(t(),n(y,{key:2,color:"orange"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1})):(t(),n(y,{key:3,color:"red"},{default:e(()=>[_(s(k.$t(`chat.dict.status.${o.value.status}`)),1)]),_:1}))]))]),_:1}),a(c,{label:"\u672C\u5730IP",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",It,s(o.value.local_ip||"-"),1))]),_:1}),a(c,{label:"\u5BA2\u6237\u7AEFIP",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",Bt,s(o.value.client_ip||"-"),1))]),_:1}),a(c,{label:"\u8FDC\u7A0BIP",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",Et,s(o.value.remote_ip||"-"),1))]),_:1}),a(c,{label:"\u8BF7\u6C42\u65F6\u95F4",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",xt,s(o.value.req_time||"-"),1))]),_:1}),a(c,{label:"\u521B\u5EFA\u65F6\u95F4",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{widths:["200px"],rows:1})]),_:1})):(t(),d("span",At,s(o.value.created_at||"-"),1))]),_:1}),a(c,{label:"\u9519\u8BEF\u4FE1\u606F",span:2},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:1})]),_:1})):(t(),d("span",St,s(o.value.err_msg||"-"),1))]),_:1})]),_:1})),[[q,["admin"]]]),h((t(),n(I,{layout:"inline-vertical",column:2,style:{"margin-top":"10px",position:"relative"}},{default:e(()=>[a(c,{span:2},{default:e(()=>[a(j,{type:"card"},{default:e(()=>[a(W,{key:"1",title:"\u6A21\u578B\u4EE3\u7406"},{default:e(()=>[r(p)?(t(),n(u,{key:0,animation:!0},{default:e(()=>[a(m,{rows:3})]),_:1})):(t(),n(L,{key:1},{default:e(()=>[o.value.model_agent?(t(),n(r(Aa),{key:0,data:o.value.model_agent,"show-length":!0},null,8,["data"])):(t(),d("span",qt,"-"))]),_:1}))]),_:1})]),_:1})]),_:1})]),_:1})),[[q,["admin"]]])])}}});const Ut=pe(Lt,[["__scopeId","data-v-f5be0e2e"]]),Nt={class:"container"},Mt={class:"action-icon"},Pt={class:"action-icon"},Tt={id:"tableSetting"},Yt={style:{"margin-right":"4px",cursor:"move"}},Ht={class:"title"},Rt={name:"ImageList"},jt=fe({...Rt,setup(oe){const E=localStorage.getItem("userRole"),p=me({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),P=$([]);E==="user"&&(async()=>{try{const{data:i}=await La();P.value=i.items}catch{}})();const z=$([]);(async()=>{try{const{data:i}=await Na();z.value=i.items}catch{}})();const R=$([]);E==="admin"&&(async()=>{try{const{data:i}=await Ma();R.value=i.items}catch{}})();const T=()=>({app_id:$(),trace_id:$(),user_id:$(),key:"",models:[],model_agents:[],total_time:$(),status:$(),req_time:[ce().format("YYYY-MM-DD 00:00:00"),ce().format("YYYY-MM-DD 23:59:59")]}),{loading:Y,setLoading:k}=ge(!0),{t:v}=Ge(),m=$([]),u=$(T()),D=$([]),c=$([]),y=$("medium"),I={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},L=me({...I}),W=te(()=>[{name:v("size.mini"),value:"mini"},{name:v("size.small"),value:"small"},{name:v("size.medium"),value:"medium"},{name:v("size.large"),value:"large"}]),j=te(()=>[{title:v(E==="admin"?"chat.columns.user_id":"chat.columns.app_id"),dataIndex:E==="admin"?"user_id":"app_id",slotName:E==="admin"?"user_id":"app_id",align:"center",width:75},{title:v("chat.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:v("chat.columns.prompt"),dataIndex:"prompt",slotName:"prompt",align:"center",ellipsis:!0,tooltip:!0},{title:v("chat.columns.images"),dataIndex:"images",slotName:"images",align:"center"},{title:v("chat.columns.total_price"),dataIndex:"total_tokens",slotName:"total_tokens",align:"center"},{title:v("chat.columns.total_time"),dataIndex:"total_time",slotName:"total_time",align:"center"},{title:v("chat.columns.internal_time"),dataIndex:"internal_time",slotName:"internal_time",align:"center"},{title:v("chat.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:v("chat.columns.req_time"),dataIndex:"req_time",slotName:"req_time",align:"center",width:132},{title:v("chat.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:75}]);E==="user"&&j.value.splice(6,1);const q=te(()=>[{label:v("chat.dict.status.1"),value:1},{label:v("chat.dict.status.2"),value:2},{label:v("chat.dict.status.-1"),value:-1}]);E==="admin"&&q.value.push({label:v("chat.dict.status.3"),value:3},{label:v("chat.dict.status.-100"),value:-100});const g=async(i={...I,...u.value})=>{k(!0);try{const{data:f}=await za(i);m.value=f.items,L.current=i.current,L.pageSize=i.pageSize,L.total=f.paging.total}catch{}finally{k(!1)}},F=()=>{g({...I,...u.value})},$e=i=>{g({...I,...u.value,current:i})},Ce=i=>{I.pageSize=i,g({...I,...u.value})};g();const Ve=()=>{u.value=T()},De=(i,f)=>{y.value=i},Fe=(i,f,x)=>{i?D.value.splice(x,0,f):D.value=c.value.filter(B=>B.dataIndex!==f.dataIndex)},ne=(i,f,x,B=!1)=>{const A=B?le(i):i;return f>-1&&x>-1&&A.splice(f,1,A.splice(x,1,A[f]).pop()),A},Ie=i=>{i&&ua(()=>{const f=document.getElementById("tableSetting");new Ua(f,{onEnd(x){const{oldIndex:B,newIndex:A}=x;ne(D.value,B,A),ne(c.value,B,A)}})})};ve(()=>j.value,i=>{D.value=le(i),D.value.forEach((f,x)=>{f.checked=!0}),c.value=le(D.value)},{deep:!0,immediate:!0});const se=$(),X=$(!1),Be=i=>{se.value=i,X.value=!0},Z=$(!1),ue=$(),Ee=i=>{Z.value=!0,ue.value=i},xe=()=>{Z.value=!1};return(i,f)=>{const x=Je,B=ia,A=da,G=ra,U=_a,V=ma,C=ca,J=pa,ie=fa,O=va,de=ya,Ae=ka,re=ha,Se=ba,K=ga,_e=Oe,qe=be,ee=wa,ze=Ke,Le=$a,Ue=Ca,Ne=Qe,Me=We,Pe=Va,Te=Da,Ye=Fa,b=he,He=Ia,Re=Ba,je=Ea,w=ye("permission");return t(),d("div",Nt,[a(A,{class:"container-breadcrumb"},{default:e(()=>[a(B,null,{default:e(()=>[a(x)]),_:1}),a(B,null,{default:e(()=>[_(s(i.$t("menu.image")),1)]),_:1}),a(B,null,{default:e(()=>[_(s(i.$t("menu.image.list")),1)]),_:1})]),_:1}),a(je,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:e(()=>[a(O,null,{default:e(()=>[a(C,{flex:1},{default:e(()=>[a(Ae,{model:u.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:e(()=>[h((t(),n(O,{gutter:16},{default:e(()=>[a(C,{span:8},{default:e(()=>[a(V,{field:"app_id",label:i.$t("chat.form.app_id")},{default:e(()=>[a(U,{modelValue:u.value.app_id,"onUpdate:modelValue":f[0]||(f[0]=l=>u.value.app_id=l),placeholder:i.$t("chat.form.selectDefault"),"allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),d(N,null,M(P.value,l=>(t(),n(G,{key:l.app_id,value:l.app_id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"key",label:i.$t("chat.form.key")},{default:e(()=>[a(J,{modelValue:u.value.key,"onUpdate:modelValue":f[1]||(f[1]=l=>u.value.key=l),placeholder:i.$t("chat.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"trace_id",label:i.$t("chat.form.user.trace_id")},{default:e(()=>[a(J,{modelValue:u.value.trace_id,"onUpdate:modelValue":f[2]||(f[2]=l=>u.value.trace_id=l),placeholder:i.$t("chat.form.trace_id.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"models",label:i.$t("chat.form.models")},{default:e(()=>[a(U,{modelValue:u.value.models,"onUpdate:modelValue":f[3]||(f[3]=l=>u.value.models=l),placeholder:i.$t("chat.form.selectDefault"),"max-tag-count":1,multiple:"","allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),d(N,null,M(z.value,l=>(t(),n(G,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"status",label:i.$t("chat.form.status")},{default:e(()=>[a(U,{modelValue:u.value.status,"onUpdate:modelValue":f[4]||(f[4]=l=>u.value.status=l),options:r(q),placeholder:i.$t("chat.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"req_time",label:i.$t("chat.form.req_time")},{default:e(()=>[a(ie,{modelValue:u.value.req_time,"onUpdate:modelValue":f[5]||(f[5]=l=>u.value.req_time=l),placeholder:["\u5F00\u59CB\u65F6\u95F4","\u7ED3\u675F\u65F6\u95F4"],"time-picker-props":{defaultValue:["00:00:00","23:59:59"]},"show-time":""},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})),[[w,["user"]]]),h((t(),n(O,{gutter:16},{default:e(()=>[a(C,{span:5},{default:e(()=>[a(V,{field:"trace_id",label:i.$t("chat.form.trace_id"),"label-col-props":{span:6}},{default:e(()=>[a(J,{modelValue:u.value.trace_id,"onUpdate:modelValue":f[6]||(f[6]=l=>u.value.trace_id=l),placeholder:i.$t("chat.form.trace_id.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:6},{default:e(()=>[a(V,{field:"models",label:i.$t("chat.form.models"),"label-col-props":{span:5}},{default:e(()=>[a(U,{modelValue:u.value.models,"onUpdate:modelValue":f[7]||(f[7]=l=>u.value.models=l),placeholder:i.$t("chat.form.selectDefault"),"max-tag-count":1,multiple:"","allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),d(N,null,M(z.value,l=>(t(),n(G,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:5},{default:e(()=>[a(V,{field:"key",label:i.$t("chat.form.key"),"label-col-props":{span:6}},{default:e(()=>[a(J,{modelValue:u.value.key,"onUpdate:modelValue":f[8]||(f[8]=l=>u.value.key=l),placeholder:i.$t("chat.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"total_time",label:i.$t("chat.form.total_time")},{default:e(()=>[a(de,{modelValue:u.value.total_time,"onUpdate:modelValue":f[9]||(f[9]=l=>u.value.total_time=l),precision:0,min:1,placeholder:i.$t("chat.form.total_time.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:5},{default:e(()=>[a(V,{field:"user_id",label:i.$t("chat.form.user_id"),"label-col-props":{span:6}},{default:e(()=>[a(de,{modelValue:u.value.user_id,"onUpdate:modelValue":f[10]||(f[10]=l=>u.value.user_id=l),placeholder:i.$t("chat.form.user_id.placeholder"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:6},{default:e(()=>[a(V,{field:"model_agents",label:i.$t("key.form.modelAgents"),"label-col-props":{span:5}},{default:e(()=>[a(U,{modelValue:u.value.model_agents,"onUpdate:modelValue":f[11]||(f[11]=l=>u.value.model_agents=l),placeholder:i.$t("key.form.selectDefault"),"max-tag-count":1,multiple:"","allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),d(N,null,M(R.value,l=>(t(),n(G,{key:l.id,value:l.id,label:l.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:5},{default:e(()=>[a(V,{field:"status",label:i.$t("chat.form.status"),"label-col-props":{span:6}},{default:e(()=>[a(U,{modelValue:u.value.status,"onUpdate:modelValue":f[12]||(f[12]=l=>u.value.status=l),options:r(q),placeholder:i.$t("chat.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(C,{span:8},{default:e(()=>[a(V,{field:"req_time",label:i.$t("chat.form.req_time")},{default:e(()=>[a(ie,{modelValue:u.value.req_time,"onUpdate:modelValue":f[13]||(f[13]=l=>u.value.req_time=l),placeholder:["\u5F00\u59CB\u65F6\u95F4","\u7ED3\u675F\u65F6\u95F4"],"time-picker-props":{defaultValue:["00:00:00","23:59:59"]},"show-time":""},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})),[[w,["admin"]]])]),_:1},8,["model"])]),_:1}),a(re,{style:{height:"84px"},direction:"vertical"}),a(C,{flex:"86px",style:{"text-align":"right"}},{default:e(()=>[a(qe,{direction:"vertical",size:18},{default:e(()=>[a(K,{type:"primary",onClick:F},{icon:e(()=>[a(Se)]),default:e(()=>[_(" "+s(i.$t("chat.form.search")),1)]),_:1}),a(K,{onClick:Ve},{icon:e(()=>[a(_e)]),default:e(()=>[_(" "+s(i.$t("chat.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(re,{style:{"margin-top":"0"}}),a(O,{style:{"margin-bottom":"16px"}},{default:e(()=>[a(C,{span:24,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:e(()=>[a(ee,{content:i.$t("actions.refresh")},{default:e(()=>[S("div",{class:"action-icon",onClick:F},[a(_e,{size:"18"})])]),_:1},8,["content"]),a(Ue,{onSelect:De},{content:e(()=>[(t(!0),d(N,null,M(r(W),l=>(t(),n(Le,{key:l.value,value:l.value,class:sa({active:l.value===y.value})},{default:e(()=>[S("span",null,s(l.name),1)]),_:2},1032,["value","class"]))),128))]),default:e(()=>[a(ee,{content:i.$t("actions.density")},{default:e(()=>[S("div",Mt,[a(ze,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(ee,{content:i.$t("actions.column_setting")},{default:e(()=>[a(Te,{trigger:"click",position:"bl",onPopupVisibleChange:Ie},{content:e(()=>[S("div",Tt,[(t(!0),d(N,null,M(c.value,(l,H)=>(t(),d("div",{key:l.dataIndex,class:"setting"},[S("div",Yt,[a(Me)]),S("div",null,[a(Pe,{modelValue:l.checked,"onUpdate:modelValue":ae=>l.checked=ae,onChange:ae=>Fe(ae,l,H)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),S("div",Ht,s(l.title==="#"?"\u5E8F\u5217\u53F7":l.title),1)]))),128))])]),default:e(()=>[S("div",Pt,[a(Ne,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(He,{"row-key":"id",loading:r(Y),pagination:L,columns:D.value,data:m.value,bordered:!1,size:y.value,"row-selection":p,onPageChange:$e,onPageSizeChange:Ce},{user_id:e(({record:l})=>[_(s(l.is_smart_match?"-":l.user_id),1)]),images:e(({record:l})=>[a(K,{type:"text",size:"small",onClick:H=>Be(l.id)},{default:e(()=>[_("\u67E5\u770B")]),_:2},1032,["onClick"]),se.value===l.id?(t(),n(Ye,{key:0,visible:X.value,"onUpdate:visible":f[14]||(f[14]=H=>X.value=H),"src-list":l.images},null,8,["visible","src-list"])):ke("",!0)]),total_tokens:e(({record:l})=>[_(s(l.total_tokens?`$${r(xa)(l.total_tokens)}`:l.status===1&&l.billing_method===2?0:"-"),1)]),total_time:e(({record:l})=>[l.total_time>18e4?h((t(),n(b,{key:0,color:"red"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["user"]]]):l.total_time>12e4?h((t(),n(b,{key:1,color:"orange"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["user"]]]):l.total_time>9e4?h((t(),n(b,{key:2,color:"gold"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["user"]]]):h((t(),n(b,{key:3,color:"green"},{default:e(()=>[_(s(l.total_time||"-"),1)]),_:2},1024)),[[w,["user"]]]),l.total_time>12e4?h((t(),n(b,{key:4,color:"red"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["admin"]]]):l.total_time>9e4?h((t(),n(b,{key:5,color:"orange"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["admin"]]]):l.total_time>6e4?h((t(),n(b,{key:6,color:"gold"},{default:e(()=>[_(s(l.total_time),1)]),_:2},1024)),[[w,["admin"]]]):h((t(),n(b,{key:7,color:"green"},{default:e(()=>[_(s(l.total_time||"-"),1)]),_:2},1024)),[[w,["admin"]]])]),internal_time:e(({record:l})=>[l.internal_time>1e3?h((t(),n(b,{key:0,color:"red"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["user"]]]):l.internal_time>500?h((t(),n(b,{key:1,color:"orange"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["user"]]]):l.internal_time>300?h((t(),n(b,{key:2,color:"gold"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["user"]]]):h((t(),n(b,{key:3,color:"green"},{default:e(()=>[_(s(l.internal_time||"-"),1)]),_:2},1024)),[[w,["user"]]]),l.internal_time>500?h((t(),n(b,{key:4,color:"red"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["admin"]]]):l.internal_time>300?h((t(),n(b,{key:5,color:"orange"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["admin"]]]):l.internal_time>100?h((t(),n(b,{key:6,color:"gold"},{default:e(()=>[_(s(l.internal_time),1)]),_:2},1024)),[[w,["admin"]]]):h((t(),n(b,{key:7,color:"green"},{default:e(()=>[_(s(l.internal_time||"-"),1)]),_:2},1024)),[[w,["admin"]]])]),status:e(({record:l})=>[l.status===-1?(t(),n(b,{key:0,color:"red"},{default:e(()=>[_(s(i.$t(`chat.dict.status.${l.status}`)),1)]),_:2},1024)):l.status===2?(t(),n(b,{key:1,color:"gold"},{default:e(()=>[_(s(i.$t(`chat.dict.status.${l.status}`)),1)]),_:2},1024)):l.status===3?(t(),n(b,{key:2,color:"orange"},{default:e(()=>[_(s(i.$t(`chat.dict.status.${l.status}`)),1)]),_:2},1024)):(t(),n(b,{key:3,color:"green"},{default:e(()=>[_(s(i.$t(`chat.dict.status.${l.status}`)),1)]),_:2},1024))]),operations:e(({record:l})=>[a(K,{type:"text",size:"small",onClick:H=>Ee(l.id)},{default:e(()=>[_(s(i.$t("chat.columns.operations.view")),1)]),_:2},1032,["onClick"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(Re,{title:i.$t("menu.image.detail"),width:700,footer:!1,visible:Z.value,"unmount-on-close":"","render-to-body":"",onCancel:xe},{default:e(()=>[a(Ut,{id:ue.value},null,8,["id"])]),_:1},8,["title","visible"])]),_:1})])}}});const Fl=pe(jt,[["__scopeId","data-v-5e2a3db0"]]);export{Fl as default};
