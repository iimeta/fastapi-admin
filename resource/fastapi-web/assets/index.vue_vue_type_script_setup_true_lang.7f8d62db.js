import{u as P}from"./index.9b75afac.js";/* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                */import{d as U,e as k,B as o,C as s,aH as t,aG as e,u as a,aD as d,aM as i,aL as b,aE as x,bJ as z,bK as O,bL as W,bM as X,bN as Y,b_ as Z,bH as ee}from"./arco.91d8d802.js";import{u as ae}from"./loading.d8a03711.js";import{p as g,q as $}from"./common.df364eef.js";import{f as te}from"./model.9407e965.js";const le={style:{margin:"10px 0 30px 10px"}},oe={key:1},ie={key:1},ue={key:1},ne={key:1},de={key:1},_e={key:1},se={key:1},re={key:1},me={key:1},ce={key:1},pe={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},fe={key:1},be={key:1},ye={key:1},ke={key:1},ge={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},ve={key:1},he={key:1},xe={key:1},we={key:1},$e={key:1},Ee={key:1},qe={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ce={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Fe={key:1},De={key:1},Be={key:1},Qe={key:1},Me={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ae={key:1},Ve={key:1},Ne={key:1},Se={name:"ModelDetail"},ta=U({...Se,props:{id:{type:String,default:""}},setup(G){const J=G,{t:u}=P(),{loading:r,setLoading:C}=ae(!0),n=k({}),F=k(!1),D=k([]),B=k(!1),Q=k([]),M=k(!1),A=k([]),q=k(!1),V=k(!1),N=k([]),S=k([]),L=k([]),T=k(!1),j=k([]),I=k(!1),H=k([]);return(async(w={id:J.id})=>{C(!0);try{const{data:y}=await te(w);n.value=y,y.type===2?(n.value.billing_method=2,B.value=!0,Q.value=y.image_quotas):y.type===5||y.type===6?(n.value.billing_method=y.audio_quota.billing_method,M.value=!0,A.value[0]=y.audio_quota):y.type===100?(q.value=!0,N.value[0]=y.multimodal_quota.text_quota,S.value=y.multimodal_quota.image_quotas,y.multimodal_quota.search_quota>0&&(V.value=!0,L.value[0]=y.multimodal_quota)):y.type===101?(T.value=!0,j.value[0]=y.realtime_quota):y.type===102?(I.value=!0,H.value[0]=y.multimodal_audio_quota):(n.value.billing_method=y.text_quota.billing_method,F.value=!0,D.value[0]=y.text_quota)}catch{}finally{C(!1)}})(),(w,y)=>{const m=z,c=O,p=W,R=X,K=Y,f=Z,h=ee;return o(),s("div",le,[t(K,{column:2,bordered:"","label-style":{padding:"5px 8px 5px 15px"},"value-style":{width:"350px",padding:"5px 8px 5px 15px"}},{default:e(()=>[t(p,{label:a(u)("common.corp")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",oe,i(n.value.corp_name),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.name")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ie,i(n.value.name),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.model")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ue,i(n.value.model),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.type")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ne,i(a(u)(`dict.model_type.${n.value.type}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.baseUrl")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",de,i(n.value.base_url||"-"),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.path")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",_e,i(n.value.path||"-"),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.dataFormat")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",se,i(a(u)(`model.dict.data_format.${n.value.data_format}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.isPublic")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",re,i(a(u)(`model.dict.is_public.${n.value.is_public}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.is_enable_preset_config")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",me,i(a(u)(`model.dict.is_enable_preset_config.${n.value.is_enable_preset_config||!1}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.preset_config.is_support_system_role")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ce,i(a(u)(`model.dict.is_support_system_role.${((_=(l=n.value)==null?void 0:l.preset_config)==null?void 0:_.is_support_system_role)||!1}`)),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.preset_config.system_role_prompt"),span:2},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",pe,i(((_=(l=n.value)==null?void 0:l.preset_config)==null?void 0:_.system_role_prompt)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.preset_config.min_tokens")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",fe,i(((_=(l=n.value)==null?void 0:l.preset_config)==null?void 0:_.min_tokens)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.preset_config.max_tokens")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",be,i(((_=(l=n.value)==null?void 0:l.preset_config)==null?void 0:_.max_tokens)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.isEnableModelAgent")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ye,i(a(u)(`model.dict.is_enable_model_agent.${n.value.is_enable_model_agent}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.lb_strategy")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ke,i(w.$t(`dict.lb_strategy.${n.value.lb_strategy||1}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.modelAgentNames")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ge,i(((_=(l=n.value)==null?void 0:l.model_agent_names)==null?void 0:_.join(`
`))||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.isEnableForward"),span:2},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",ve,i(a(u)(`model.dict.is_enable_forward.${n.value.is_enable_forward||!1}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.forwardRule")},{default:e(()=>{var l,_,v,E;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",he,i((_=(l=n.value)==null?void 0:l.forward_config)!=null&&_.forward_rule?a(u)(`model.dict.forward_rule.${((E=(v=n.value)==null?void 0:v.forward_config)==null?void 0:E.forward_rule)||1}`):"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.content_length")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",xe,i(((_=(l=n.value)==null?void 0:l.forward_config)==null?void 0:_.content_length)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.targetModelName")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",we,i(((_=(l=n.value)==null?void 0:l.forward_config)==null?void 0:_.target_model_name)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.matchRule")},{default:e(()=>{var l,_,v,E;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",$e,i((_=(l=n.value)==null?void 0:l.forward_config)!=null&&_.match_rule?a(u)(`model.dict.match_rule.${((E=(v=n.value)==null?void 0:v.forward_config)==null?void 0:E.match_rule)||1}`):"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.decisionModelName")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Ee,i(((_=(l=n.value)==null?void 0:l.forward_config)==null?void 0:_.decision_model_name)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.keywords")},{default:e(()=>{var l,_,v;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",qe,i(((v=(_=(l=n.value)==null?void 0:l.forward_config)==null?void 0:_.keywords)==null?void 0:v.join(`
`))||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.targetModelNames")},{default:e(()=>{var l,_,v;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Ce,i(((v=(_=(l=n.value)==null?void 0:l.forward_config)==null?void 0:_.target_model_names)==null?void 0:v.join(`
`))||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.is_enable_fallback")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Fe,i(a(u)(`model.dict.is_enable_fallback.${n.value.is_enable_fallback||!1}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.fallback_model_agent_name")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",De,i(((_=(l=n.value)==null?void 0:l.fallback_config)==null?void 0:_.model_agent_name)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.fallback_model_name")},{default:e(()=>{var l,_;return[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Be,i(((_=(l=n.value)==null?void 0:l.fallback_config)==null?void 0:_.model_name)||"-"),1))]}),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.billingMethod")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Qe,i(w.$t(`model.dict.billing_method.${n.value.billing_method||1}`)),1))]),_:1},8,["label"]),t(p,{label:a(u)("model.detail.label.remark")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Me,i(n.value.remark||"-"),1))]),_:1},8,["label"]),t(p,{label:a(u)("common.status")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Ae,[n.value.status===1?(o(),d(R,{key:0,color:"green"},{default:e(()=>[b(i(w.$t(`dict.status.${n.value.status}`)),1)]),_:1})):(o(),d(R,{key:1,color:"red"},{default:e(()=>[b(i(w.$t(`dict.status.${n.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),t(p,{label:a(u)("common.created_at")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Ve,i(n.value.created_at),1))]),_:1},8,["label"]),t(p,{label:a(u)("common.updated_at")},{default:e(()=>[a(r)?(o(),d(c,{key:0,animation:!0},{default:e(()=>[t(m,{rows:1})]),_:1})):(o(),s("span",Ne,i(n.value.updated_at),1))]),_:1},8,["label"])]),_:1}),F.value?(o(),d(h,{key:0,style:{"margin-top":"15px"},data:D.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(l.billing_method===1?`$${a(g)(l.prompt_ratio)}/k`:"-"),1)]),_:1}),t(f,{title:"\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(l.billing_method===1?`$${a(g)(l.completion_ratio)}/k`:`$${a($)(l.fixed_quota)}/\u6B21`),1)]),_:1})]),_:1},8,["data"])):x("",!0),B.value?(o(),d(h,{key:1,style:{"margin-top":"15px"},data:Q.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u5BBD\u5EA6","data-index":"width",align:"center"}),t(f,{title:"\u9AD8\u5EA6","data-index":"height",align:"center"}),t(f,{title:"\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:e(({record:l})=>[b(i(`$${a($)(l.fixed_quota)}/\u5F20`),1)]),_:1}),t(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:e(({record:l})=>[b(i(l.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])):x("",!0),M.value?(o(),d(h,{key:2,style:{"margin-top":"15px"},data:A.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(n.value.type===5?l.billing_method===1?`$${a(g)(l.prompt_ratio)}/k`:`$${a($)(l.fixed_quota)}/\u6B21`:"-"),1)]),_:1}),t(f,{title:"\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(n.value.type===6?l.billing_method===1?`$${a(g)(l.completion_ratio)}/min`:`$${a($)(l.fixed_quota)}/\u6B21`:"-"),1)]),_:1})]),_:1},8,["data"])):x("",!0),q.value?(o(),d(h,{key:3,style:{"margin-top":"15px"},data:N.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.prompt_ratio)}/k`),1)]),_:1}),t(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):x("",!0),q.value?(o(),d(h,{key:4,style:{"margin-top":"15px"},data:S.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u8BC6\u56FE\u6A21\u5F0F","data-index":"mode",align:"center"}),t(f,{title:"\u8BC6\u56FE\u4EF7\u683C","data-index":"fixed_quota",align:"center"},{cell:e(({record:l})=>[b(i(`$${a($)(l.fixed_quota)}/\u5F20`),1)]),_:1}),t(f,{title:"\u9ED8\u8BA4","data-index":"is_default",align:"center"},{cell:e(({record:l})=>[b(i(l.is_default?"\u662F":"-"),1)]),_:1})]),_:1},8,["data"])):x("",!0),V.value?(o(),d(h,{key:5,style:{"margin-top":"15px"},data:L.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u641C\u7D22\u4EF7\u683C","data-index":"search_quota",align:"center"},{cell:e(({record:l})=>[b(i(`$${a($)(l.search_quota)}/\u6B21`),1)]),_:1})]),_:1},8,["data"])):x("",!0),T.value?(o(),d(h,{key:6,style:{"margin-top":"15px"},data:j.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.text_quota.prompt_ratio)}/k`),1)]),_:1}),t(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.text_quota.completion_ratio)}/k`),1)]),_:1}),t(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.audio_quota.prompt_ratio)}/k`),1)]),_:1}),t(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):x("",!0),I.value?(o(),d(h,{key:7,style:{"margin-top":"15px"},data:H.value,pagination:!1,bordered:!1},{columns:e(()=>[t(f,{title:"\u6587\u672C\u63D0\u95EE\u4EF7\u683C","data-index":"text_quota.prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.text_quota.prompt_ratio)}/k`),1)]),_:1}),t(f,{title:"\u6587\u672C\u56DE\u7B54\u4EF7\u683C","data-index":"text_quota.completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.text_quota.completion_ratio)}/k`),1)]),_:1}),t(f,{title:"\u97F3\u9891\u63D0\u95EE\u4EF7\u683C","data-index":"audio_quota.prompt_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.audio_quota.prompt_ratio)}/k`),1)]),_:1}),t(f,{title:"\u97F3\u9891\u56DE\u7B54\u4EF7\u683C","data-index":"audio_quota.completion_ratio",align:"center"},{cell:e(({record:l})=>[b(i(`$${a(g)(l.audio_quota.completion_ratio)}/k`),1)]),_:1})]),_:1},8,["data"])):x("",!0)])}}});export{ta as _};
