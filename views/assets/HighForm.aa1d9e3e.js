import{u as T,R as v,j as e,a as i,c as t,t as m,C as x,b as y,an as A,ao as k}from"./index.2ebf406f.js";import{B as p}from"./index.f3dd4729.js";import{F as l,I as h,S as u,T as F}from"./index.a67bce5b.js";import"./index.271b39e3.js";const I=()=>{const{config:f}=T(),{tsn:g}=f,[d]=l.useForm(),[s,D]=v.useState(void 0),b=()=>{d.validateFields().then(async c=>{const{ip:r}=await A(),{code:n,msg:a}=await k({...c,ip:r,tsn:g});D({code:n,msg:a})})},C=()=>{};return e("div",{className:"w-[572px] mx-auto",children:i(l,{form:d,colon:!1,labelCol:{span:5},requiredMark:!1,css:t`
                        .ant-input-clear-icon,
                        .ant-form-item-feedback-icon { display: inline-flex;};
                    `,children:[e(l.Item,{label:"T-One\u57DF\u540D",name:"domain",hasFeedback:!0,validateTrigger:["onBlur"],rules:[{required:!0,message:"T-One\u57DF\u540D\u57DF\u540D\u4E0D\u80FD\u4E3A\u7A7A"},{async validator(c,r,n){if(!r)return Promise.reject();const{code:a,msg:o}=await m({url:r});return a!==200?Promise.reject(o):Promise.resolve()}}],className:"!mb-3",children:e(h,{placeholder:"\u8BF7\u8F93\u5165T-One\u57DF\u540D"})}),e(l.Item,{label:"TestLib\u57DF\u540D",name:"testlib",hasFeedback:!0,validateTrigger:["onBlur"],rules:[{required:!0,message:"T-TestLib\u57DF\u540D\u4E0D\u80FD\u4E3A\u7A7A"},{async validator(c,r,n){if(!r)return Promise.reject();const{code:a,msg:o}=await m({url:r});return a!==200?Promise.reject(o):Promise.resolve()}}],help:i(u,{direction:"vertical",size:0,children:[s&&s.code===200&&i(u,{align:"center",css:t`height: 32px;`,children:[e(x,{css:t`display: inline-flex;svg{ fill: #52c41a;}`}),e(F.Text,{style:{color:"rgba(0,0,0,.65)"},children:"\u673A\u5668\u6CE8\u518C\u6210\u529F"})]}),s&&s.code!==200&&i(u,{align:"center",css:t`height: 32px;`,children:[e(y,{css:t`display: inline-flex;svg{ fill: #ff4d4f;}`}),e(F.Text,{type:"danger",children:s.msg||"\u673A\u5668\u6CE8\u518C\u5931\u8D25"})]})]}),children:e(h,{placeholder:"\u8BF7\u8F93\u5165TestLib\u57DF\u540D"})}),e(l.Item,{label:" ",children:i(u,{children:[e(p,{type:"primary",onClick:b,children:"\u6CE8\u518C\u673A\u5668"}),e(p,{onClick:C,children:"\u767B\u5F55TestLib"})]})})]})})};export{I as default};
