import{u as J}from"./index.a8f8f038.js";/* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                */import{d as K,e as k,B as o,C as s,aH as l,aG as a,u as e,aD as _,aM as i,aL as b,aE as h,bJ as P,bK as U,bL as z,bM as O,bN as W,b_ as X,bH as Y}from"./arco.553b67be.js";import{u as Z}from"./loading.9398a00f.js";import{p as g,q as $}from"./common.df364eef.js";import{f as ee}from"./model.78bc160f.js";const ae={style:{margin:"10px 0 30px 10px"}},te={key:1},le={key:1},oe={key:1},ie={key:1},ue={key:1},ne={key:1},de={key:1},_e={key:1},se={key:1},re={key:1},me={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},ce={key:1},pe={key:1},fe={key:1},be={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},ye={key:1},ke={key:1},ge={key:1},ve={key:1},xe={key:1},he={key:1},we={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},$e={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ee={key:1},qe={key:1},Ce={key:1},Fe={key:1},De={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Be={key:1},Qe={key:1},Ae={name:"ModelDetail"},Ye=K({...Ae,props:{id:{type:String,default:""}},setup(H){const R=H,{t:n}=J(),{loading:r,setLoading:C}=Z(!0),d=k({}),F=k(!1),D=k([]),B=k(!1),Q=k([]),A=k(!1),M=k([]),q=k(!1),V=k([]),N=k([]),L=k(!1),T=k([]),S=k(!1),j=k([]);return(async(E={id:R.id})=>{C(!0);try{const{data:y}=await ee(E);d.value=y,y.type===2?(d.value.billing_method=2,B.value=!0,Q.value=y.image_quotas):y.type===5||y.type===6?(d.value.billing_method=y.audio_quota.billing_method,A.value=!0,M.value[0]=y.audio_quota):y.type===100?(q.value=!0,V.value[0]=y.multimodal_quota.text_quota,N.value=y.multimodal_quota.image_quotas):y.type===101?(L.value=!0,T.value[0]=y.realtime_quota):y.type===102?(S.value=!0,j.value[0]=y.multimodal_audio_quota):(d.value.billing_method=y.text_quota.billing_method,F.value=!0,D.value[0]=y.text_quota)}catch{}finally{C(!1)}})(),(E,y)=>{const m=P,c=U,p=z,I=O,G=W,f=X,x=Y;return o(),s("div",ae,[l(G,{column:2,bordered:"","label-style":{padding:"5px 8px 5px 15px"},"value-style":{width:"350px",padding:"5px 8px 5px 15px"}},{default:a(()=>[l(p,{label:e(n)("common.corp")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",te,i(d.value.corp_name),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.name")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",le,i(d.value.name),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.model")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",oe,i(d.value.model),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.type")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ie,i(e(n)(`dict.model_type.${d.value.type}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.baseUrl")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ue,i(d.value.base_url||"-"),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.path")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ne,i(d.value.path||"-"),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.dataFormat")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",de,i(e(n)(`model.dict.data_format.${d.value.data_format}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.isPublic")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",_e,i(e(n)(`model.dict.is_public.${d.value.is_public}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.is_enable_preset_config")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",se,i(e(n)(`model.dict.is_enable_preset_config.${d.value.is_enable_preset_config||!1}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.preset_config.is_support_system_role")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",re,i(e(n)(`model.dict.is_support_system_role.${((u=(t=d.value)==null?void 0:t.preset_config)==null?void 0:u.is_support_system_role)||!1}`)),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.preset_config.system_role_prompt"),span:2},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",me,i(((u=(t=d.value)==null?void 0:t.preset_config)==null?void 0:u.system_role_prompt)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.preset_config.min_tokens")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ce,i(((u=(t=d.value)==null?void 0:t.preset_config)==null?void 0:u.min_tokens)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.preset_config.max_tokens")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",pe,i(((u=(t=d.value)==null?void 0:t.preset_config)==null?void 0:u.max_tokens)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.isEnableModelAgent")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",fe,i(e(n)(`model.dict.is_enable_model_agent.${d.value.is_enable_model_agent}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.modelAgentNames")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",be,i(((u=(t=d.value)==null?void 0:t.model_agent_names)==null?void 0:u.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.isEnableForward")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ye,i(e(n)(`model.dict.is_enable_forward.${d.value.is_enable_forward||!1}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.forwardRule")},{default:a(()=>{var t,u,v,w;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ke,i((u=(t=d.value)==null?void 0:t.forward_config)!=null&&u.forward_rule?e(n)(`model.dict.forward_rule.${((w=(v=d.value)==null?void 0:v.forward_config)==null?void 0:w.forward_rule)||1}`):"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.content_length")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ge,i(((u=(t=d.value)==null?void 0:t.forward_config)==null?void 0:u.content_length)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.targetModelName")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",ve,i(((u=(t=d.value)==null?void 0:t.forward_config)==null?void 0:u.target_model_name)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.matchRule")},{default:a(()=>{var t,u,v,w;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",xe,i((u=(t=d.value)==null?void 0:t.forward_config)!=null&&u.match_rule?e(n)(`model.dict.match_rule.${((w=(v=d.value)==null?void 0:v.forward_config)==null?void 0:w.match_rule)||1}`):"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.decisionModelName")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",he,i(((u=(t=d.value)==null?void 0:t.forward_config)==null?void 0:u.decision_model_name)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.keywords")},{default:a(()=>{var t,u,v;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",we,i(((v=(u=(t=d.value)==null?void 0:t.forward_config)==null?void 0:u.keywords)==null?void 0:v.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.targetModelNames")},{default:a(()=>{var t,u,v;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",$e,i(((v=(u=(t=d.value)==null?void 0:t.forward_config)==null?void 0:u.target_model_names)==null?void 0:v.join(`
`))||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.is_enable_fallback")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",Ee,i(e(n)(`model.dict.is_enable_fallback.${d.value.is_enable_fallback||!1}`)),1))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.fallback_model_agent_name")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",qe,i(((u=(t=d.value)==null?void 0:t.fallback_config)==null?void 0:u.model_agent_name)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.fallback_model_name")},{default:a(()=>{var t,u;return[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",Ce,i(((u=(t=d.value)==null?void 0:t.fallback_config)==null?void 0:u.model_name)||"-"),1))]}),_:1},8,["label"]),l(p,{label:e(n)("common.status")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",Fe,[d.value.status===1?(o(),_(I,{key:0,color:"green"},{default:a(()=>[b(i(E.$t(`dict.status.${d.value.status}`)),1)]),_:1})):(o(),_(I,{key:1,color:"red"},{default:a(()=>[b(i(E.$t(`dict.status.${d.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),l(p,{label:e(n)("model.detail.label.remark")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",De,i(d.value.remark||"-"),1))]),_:1},8,["label"]),l(p,{label:e(n)("common.created_at"),span:2},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",Be,i(d.value.created_at),1))]),_:1},8,["label"]),l(p,{label:e(n)("common.updated_at")},{default:a(()=>[e(r)?(o(),_(c,{key:0,animation:!0},{default:a(()=>[l(m,{rows:1})]),_:1})):(o(),s("span",Qe,i(d.value.updated_at),1))]),_:1},8,["label"])]),_:1}),F.value?(o(),_(x,{key:0,style:{"margin-top":"15px"},data:D.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(t.billing_method===1?`$${e(g)(t.prompt_ratio)}/k`:"-"),1)]),_:1}),l(f,{title:"\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(t.billing_method===1?`$${e(g)(t.completion_ratio)}/k`:`$${e($)(t.fixed_quota)}/\u6B21`),1)]),_:1})]),_:1},8,["data"])):h("",!0),B.value?(o(),_(x,{key:1,style:{"margin-top":"15px"},data:Q.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u5BBD\u5EA6","data-index":"width",align:"center"}),l(f,{title:"\u9AD8\u5EA6","data-index":"height",align:"center"}),l(f,{title:"\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:a(({record:t})=>[b(i(`$${e($)(t.fixed_quota)}/\u5F20`),1)]),_:1}),l(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:a(({record:t})=>[b(i(t.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])):h("",!0),A.value?(o(),_(x,{key:2,style:{"margin-top":"15px"},data:M.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(d.value.type===5?t.billing_method===1?`$${e(g)(t.prompt_ratio)}/k`:`$${e($)(t.fixed_quota)}/\u6B21`:"-"),1)]),_:1}),l(f,{title:"\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(d.value.type===6?t.billing_method===1?`$${e(g)(t.completion_ratio)}/min`:`$${e($)(t.fixed_quota)}/\u6B21`:"-"),1)]),_:1})]),_:1},8,["data"])):h("",!0),q.value?(o(),_(x,{key:3,style:{"margin-top":"15px"},data:V.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.prompt_ratio)}/k`),1)]),_:1}),l(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):h("",!0),q.value?(o(),_(x,{key:4,style:{"margin-top":"15px"},data:N.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u8BC6\u56FE\u6A21\u5F0F","data-index":"mode",align:"center"}),l(f,{title:"\u8BC6\u56FE\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:a(({record:t})=>[b(i(`$${e($)(t.fixed_quota)}/\u5F20`),1)]),_:1}),l(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:a(({record:t})=>[b(i(t.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])):h("",!0),L.value?(o(),_(x,{key:5,style:{"margin-top":"15px"},data:T.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.text_quota.prompt_ratio)}/k`),1)]),_:1}),l(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.text_quota.completion_ratio)}/k`),1)]),_:1}),l(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.audio_quota.prompt_ratio)}/k`),1)]),_:1}),l(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):h("",!0),S.value?(o(),_(x,{key:6,style:{"margin-top":"15px"},data:j.value,pagination:!1,bordered:!1},{columns:a(()=>[l(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.text_quota.prompt_ratio)}/k`),1)]),_:1}),l(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.text_quota.completion_ratio)}/k`),1)]),_:1}),l(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.audio_quota.prompt_ratio)}/k`),1)]),_:1}),l(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:a(({record:t})=>[b(i(`$${e(g)(t.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):h("",!0)])}}});export{Ye as _};