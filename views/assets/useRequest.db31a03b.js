import{r as v,am as q,aa as K}from"./index.8135c6cc.js";var Q=function(e){return function(r,t){var a=v.exports.useRef(!1);e(function(){return function(){a.current=!1}},[]),e(function(){if(!a.current)a.current=!0;else return r()},t)}},Y=function(e){return typeof e=="function"};function P(n){var e=v.exports.useRef(n);e.current=v.exports.useMemo(function(){return n},[n]);var r=v.exports.useRef();return r.current||(r.current=function(){for(var t=[],a=0;a<arguments.length;a++)t[a]=arguments[a];return e.current.apply(this,t)}),r.current}const M=Q(v.exports.useEffect);var Z=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},N=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(Z(arguments[e]));return n},V=function(e,r){var t=r.manual,a=r.ready,i=a===void 0?!0:a,u=r.defaultParams,s=u===void 0?[]:u,f=r.refreshDeps,o=f===void 0?[]:f,l=r.refreshDepsAction,c=v.exports.useRef(!1);return c.current=!1,M(function(){!t&&i&&(c.current=!0,e.run.apply(e,N(s)))},[i]),M(function(){c.current||t||(c.current=!0,l?l():e.refresh())},N(o)),{onBefore:function(){if(!i)return{stopNow:!0}}}};V.onInit=function(n){var e=n.ready,r=e===void 0?!0:e,t=n.manual;return{loading:!t&&r}};const k=V;function I(n,e){if(n===e)return!0;for(var r=0;r<n.length;r++)if(!Object.is(n[r],e[r]))return!1;return!0}function X(n,e){var r=v.exports.useRef({deps:e,obj:void 0,initialized:!1}).current;return(r.initialized===!1||!I(r.deps,e))&&(r.deps=e,r.obj=n(),r.initialized=!0),r.obj}function J(n){var e=v.exports.useRef(n);return e.current=n,e}var ee=function(e){var r=J(e);v.exports.useEffect(function(){return function(){r.current()}},[])};const L=ee;var F=globalThis&&globalThis.__assign||function(){return F=Object.assign||function(n){for(var e,r=1,t=arguments.length;r<t;r++){e=arguments[r];for(var a in e)Object.prototype.hasOwnProperty.call(e,a)&&(n[a]=e[a])}return n},F.apply(this,arguments)},E=new Map,re=function(e,r,t){var a=E.get(e);a!=null&&a.timer&&clearTimeout(a.timer);var i=void 0;r>-1&&(i=setTimeout(function(){E.delete(e)},r)),E.set(e,F(F({},t),{timer:i}))},ne=function(e){return E.get(e)},j=new Map,te=function(e){return j.get(e)},ae=function(e,r){j.set(e,r),r.then(function(t){return j.delete(e),t}).catch(function(){j.delete(e)})},T={},ie=function(e,r){T[e]&&T[e].forEach(function(t){return t(r)})},D=function(e,r){return T[e]||(T[e]=[]),T[e].push(r),function(){var a=T[e].indexOf(r);T[e].splice(a,1)}},oe=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},ue=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(oe(arguments[e]));return n},le=function(e,r){var t=r.cacheKey,a=r.cacheTime,i=a===void 0?5*60*1e3:a,u=r.staleTime,s=u===void 0?0:u,f=r.setCache,o=r.getCache,l=v.exports.useRef(),c=v.exports.useRef(),g=function(m,d){f?f(d):re(m,i,d),ie(m,d.data)},p=function(m,d){return d===void 0&&(d=[]),o?o(d):ne(m)};return X(function(){if(!!t){var h=p(t);h&&Object.hasOwnProperty.call(h,"data")&&(e.state.data=h.data,e.state.params=h.params,(s===-1||new Date().getTime()-h.time<=s)&&(e.state.loading=!1)),l.current=D(t,function(m){e.setState({data:m})})}},[]),L(function(){var h;(h=l.current)===null||h===void 0||h.call(l)}),t?{onBefore:function(m){var d=p(t,m);return!d||!Object.hasOwnProperty.call(d,"data")?{}:s===-1||new Date().getTime()-d.time<=s?{loading:!1,data:d==null?void 0:d.data,returnNow:!0}:{data:d==null?void 0:d.data}},onRequest:function(m,d){var b=te(t);return b&&b!==c.current?{servicePromise:b}:(b=m.apply(void 0,ue(d)),c.current=b,ae(t,b),{servicePromise:b})},onSuccess:function(m,d){var b;t&&((b=l.current)===null||b===void 0||b.call(l),g(t,{data:m,params:d,time:new Date().getTime()}),l.current=D(t,function(C){e.setState({data:C})}))},onMutate:function(m){var d;t&&((d=l.current)===null||d===void 0||d.call(l),g(t,{data:m,params:e.state.params,time:new Date().getTime()}),l.current=D(t,function(b){e.setState({data:b})}))}}:{}};const se=le;var ce=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},fe=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(ce(arguments[e]));return n},de=function(e,r){var t=r.debounceWait,a=r.debounceLeading,i=r.debounceTrailing,u=r.debounceMaxWait,s=v.exports.useRef(),f=v.exports.useMemo(function(){var o={};return a!==void 0&&(o.leading=a),i!==void 0&&(o.trailing=i),u!==void 0&&(o.maxWait=u),o},[a,i,u]);return v.exports.useEffect(function(){if(t){var o=e.runAsync.bind(e);return s.current=q(function(l){l()},t,f),e.runAsync=function(){for(var l=[],c=0;c<arguments.length;c++)l[c]=arguments[c];return new Promise(function(g,p){var h;(h=s.current)===null||h===void 0||h.call(s,function(){o.apply(void 0,fe(l)).then(g).catch(p)})})},function(){var l;(l=s.current)===null||l===void 0||l.cancel(),e.runAsync=o}}},[t,f]),t?{onCancel:function(){var l;(l=s.current)===null||l===void 0||l.cancel()}}:{}};const ve=de;var he=function(e,r){var t=r.loadingDelay,a=v.exports.useRef();if(!t)return{};var i=function(){a.current&&clearTimeout(a.current)};return{onBefore:function(){return i(),a.current=setTimeout(function(){e.setState({loading:!0})},t),{loading:!1}},onFinally:function(){i()},onCancel:function(){i()}}};const ge=he;var pe=!!(typeof window<"u"&&window.document&&window.document.createElement);const A=pe;function H(){return A?document.visibilityState!=="hidden":!0}var O=[];function me(n){return O.push(n),function(){var r=O.indexOf(n);O.splice(r,1)}}if(A){var be=function(){if(!!H())for(var e=0;e<O.length;e++){var r=O[e];r()}};window.addEventListener("visibilitychange",be,!1)}var ye=function(e,r){var t=r.pollingInterval,a=r.pollingWhenHidden,i=a===void 0?!0:a,u=r.pollingErrorRetryCount,s=u===void 0?-1:u,f=v.exports.useRef(),o=v.exports.useRef(),l=v.exports.useRef(0),c=function(){var p;f.current&&clearTimeout(f.current),(p=o.current)===null||p===void 0||p.call(o)};return M(function(){t||c()},[t]),t?{onBefore:function(){c()},onError:function(){l.current+=1},onSuccess:function(){l.current=0},onFinally:function(){s===-1||s!==-1&&l.current<=s?f.current=setTimeout(function(){!i&&!H()?o.current=me(function(){e.refresh()}):e.refresh()},t):l.current=0},onCancel:function(){c()}}:{}};const _e=ye;var we=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},Te=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(we(arguments[e]));return n};function Pe(n,e){var r=!1;return function(){for(var t=[],a=0;a<arguments.length;a++)t[a]=arguments[a];r||(r=!0,n.apply(void 0,Te(t)),setTimeout(function(){r=!1},e))}}function xe(){return A&&typeof navigator.onLine<"u"?navigator.onLine:!0}var S=[];function Re(n){return S.push(n),function(){var r=S.indexOf(n);S.splice(r,1)}}if(A){var z=function(){if(!(!H()||!xe()))for(var e=0;e<S.length;e++){var r=S[e];r()}};window.addEventListener("visibilitychange",z,!1),window.addEventListener("focus",z,!1)}var Oe=function(e,r){var t=r.refreshOnWindowFocus,a=r.focusTimespan,i=a===void 0?5e3:a,u=v.exports.useRef(),s=function(){var o;(o=u.current)===null||o===void 0||o.call(u)};return v.exports.useEffect(function(){if(t){var f=Pe(e.refresh.bind(e),i);u.current=Re(function(){f()})}return function(){s()}},[t,i]),L(function(){s()}),{}};const Se=Oe;var Ce=function(e,r){var t=r.retryInterval,a=r.retryCount,i=v.exports.useRef(),u=v.exports.useRef(0),s=v.exports.useRef(!1);return a?{onBefore:function(){s.current||(u.current=0),s.current=!1,i.current&&clearTimeout(i.current)},onSuccess:function(){u.current=0},onError:function(){if(u.current+=1,a===-1||u.current<=a){var o=t!=null?t:Math.min(1e3*Math.pow(2,u.current),3e4);i.current=setTimeout(function(){s.current=!0,e.refresh()},o)}else u.current=0},onCancel:function(){u.current=0,i.current&&clearTimeout(i.current)}}:{}};const $e=Ce;var Ee=q,je=K,Fe="Expected a function";function Ae(n,e,r){var t=!0,a=!0;if(typeof n!="function")throw new TypeError(Fe);return je(r)&&(t="leading"in r?!!r.leading:t,a="trailing"in r?!!r.trailing:a),Ee(n,e,{leading:t,maxWait:e,trailing:a})}var Be=Ae,De=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},Me=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(De(arguments[e]));return n},We=function(e,r){var t=r.throttleWait,a=r.throttleLeading,i=r.throttleTrailing,u=v.exports.useRef(),s={};return a!==void 0&&(s.leading=a),i!==void 0&&(s.trailing=i),v.exports.useEffect(function(){if(t){var f=e.runAsync.bind(e);return u.current=Be(function(o){o()},t,s),e.runAsync=function(){for(var o=[],l=0;l<arguments.length;l++)o[l]=arguments[l];return new Promise(function(c,g){var p;(p=u.current)===null||p===void 0||p.call(u,function(){f.apply(void 0,Me(o)).then(c).catch(g)})})},function(){var o;e.runAsync=f,(o=u.current)===null||o===void 0||o.cancel()}}},[t,a,i]),t?{onCancel:function(){var o;(o=u.current)===null||o===void 0||o.cancel()}}:{}};const Le=We;var He=function(e){v.exports.useEffect(function(){e==null||e()},[])};const Ue=He;var Ne=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},ze=function(){var e=Ne(v.exports.useState({}),2),r=e[1];return v.exports.useCallback(function(){return r({})},[])};const Ge=ze;var y=globalThis&&globalThis.__assign||function(){return y=Object.assign||function(n){for(var e,r=1,t=arguments.length;r<t;r++){e=arguments[r];for(var a in e)Object.prototype.hasOwnProperty.call(e,a)&&(n[a]=e[a])}return n},y.apply(this,arguments)},qe=globalThis&&globalThis.__awaiter||function(n,e,r,t){function a(i){return i instanceof r?i:new r(function(u){u(i)})}return new(r||(r=Promise))(function(i,u){function s(l){try{o(t.next(l))}catch(c){u(c)}}function f(l){try{o(t.throw(l))}catch(c){u(c)}}function o(l){l.done?i(l.value):a(l.value).then(s,f)}o((t=t.apply(n,e||[])).next())})},Ve=globalThis&&globalThis.__generator||function(n,e){var r={label:0,sent:function(){if(i[0]&1)throw i[1];return i[1]},trys:[],ops:[]},t,a,i,u;return u={next:s(0),throw:s(1),return:s(2)},typeof Symbol=="function"&&(u[Symbol.iterator]=function(){return this}),u;function s(o){return function(l){return f([o,l])}}function f(o){if(t)throw new TypeError("Generator is already executing.");for(;r;)try{if(t=1,a&&(i=o[0]&2?a.return:o[0]?a.throw||((i=a.return)&&i.call(a),0):a.next)&&!(i=i.call(a,o[1])).done)return i;switch(a=0,i&&(o=[o[0]&2,i.value]),o[0]){case 0:case 1:i=o;break;case 4:return r.label++,{value:o[1],done:!1};case 5:r.label++,a=o[1],o=[0];continue;case 7:o=r.ops.pop(),r.trys.pop();continue;default:if(i=r.trys,!(i=i.length>0&&i[i.length-1])&&(o[0]===6||o[0]===2)){r=0;continue}if(o[0]===3&&(!i||o[1]>i[0]&&o[1]<i[3])){r.label=o[1];break}if(o[0]===6&&r.label<i[1]){r.label=i[1],i=o;break}if(i&&r.label<i[2]){r.label=i[2],r.ops.push(o);break}i[2]&&r.ops.pop(),r.trys.pop();continue}o=e.call(n,r)}catch(l){o=[6,l],a=0}finally{t=i=0}if(o[0]&5)throw o[1];return{value:o[0]?o[1]:void 0,done:!0}}},Xe=globalThis&&globalThis.__rest||function(n,e){var r={};for(var t in n)Object.prototype.hasOwnProperty.call(n,t)&&e.indexOf(t)<0&&(r[t]=n[t]);if(n!=null&&typeof Object.getOwnPropertySymbols=="function")for(var a=0,t=Object.getOwnPropertySymbols(n);a<t.length;a++)e.indexOf(t[a])<0&&Object.prototype.propertyIsEnumerable.call(n,t[a])&&(r[t[a]]=n[t[a]]);return r},Je=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},x=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(Je(arguments[e]));return n},Ke=function(){function n(e,r,t,a){a===void 0&&(a={}),this.serviceRef=e,this.options=r,this.subscribe=t,this.initState=a,this.count=0,this.state={loading:!1,params:void 0,data:void 0,error:void 0},this.state=y(y(y({},this.state),{loading:!r.manual}),a)}return n.prototype.setState=function(e){e===void 0&&(e={}),this.state=y(y({},this.state),e),this.subscribe()},n.prototype.runPluginHandler=function(e){for(var r=[],t=1;t<arguments.length;t++)r[t-1]=arguments[t];var a=this.pluginImpls.map(function(i){var u;return(u=i[e])===null||u===void 0?void 0:u.call.apply(u,x([i],r))}).filter(Boolean);return Object.assign.apply(Object,x([{}],a))},n.prototype.runAsync=function(){for(var e,r,t,a,i,u,s,f,o,l,c=[],g=0;g<arguments.length;g++)c[g]=arguments[g];return qe(this,void 0,void 0,function(){var p,h,m,d,b,C,B,$,_,w,U;return Ve(this,function(R){switch(R.label){case 0:if(this.count+=1,p=this.count,h=this.runPluginHandler("onBefore",c),m=h.stopNow,d=m===void 0?!1:m,b=h.returnNow,C=b===void 0?!1:b,B=Xe(h,["stopNow","returnNow"]),d)return[2,new Promise(function(){})];if(this.setState(y({loading:!0,params:c},B)),C)return[2,Promise.resolve(B.data)];(r=(e=this.options).onBefore)===null||r===void 0||r.call(e,c),R.label=1;case 1:return R.trys.push([1,3,,4]),$=this.runPluginHandler("onRequest",this.serviceRef.current,c).servicePromise,$||($=(U=this.serviceRef).current.apply(U,x(c))),[4,$];case 2:return _=R.sent(),p!==this.count?[2,new Promise(function(){})]:(this.setState({data:_,error:void 0,loading:!1}),(a=(t=this.options).onSuccess)===null||a===void 0||a.call(t,_,c),this.runPluginHandler("onSuccess",_,c),(u=(i=this.options).onFinally)===null||u===void 0||u.call(i,c,_,void 0),p===this.count&&this.runPluginHandler("onFinally",c,_,void 0),[2,_]);case 3:if(w=R.sent(),p!==this.count)return[2,new Promise(function(){})];throw this.setState({error:w,loading:!1}),(f=(s=this.options).onError)===null||f===void 0||f.call(s,w,c),this.runPluginHandler("onError",w,c),(l=(o=this.options).onFinally)===null||l===void 0||l.call(o,c,void 0,w),p===this.count&&this.runPluginHandler("onFinally",c,void 0,w),w;case 4:return[2]}})})},n.prototype.run=function(){for(var e=this,r=[],t=0;t<arguments.length;t++)r[t]=arguments[t];this.runAsync.apply(this,x(r)).catch(function(a){e.options.onError||console.error(a)})},n.prototype.cancel=function(){this.count+=1,this.setState({loading:!1}),this.runPluginHandler("onCancel")},n.prototype.refresh=function(){this.run.apply(this,x(this.state.params||[]))},n.prototype.refreshAsync=function(){return this.runAsync.apply(this,x(this.state.params||[]))},n.prototype.mutate=function(e){var r;Y(e)?r=e(this.state.data):r=e,this.runPluginHandler("onMutate",r),this.setState({data:r})},n}();const Qe=Ke;var W=globalThis&&globalThis.__assign||function(){return W=Object.assign||function(n){for(var e,r=1,t=arguments.length;r<t;r++){e=arguments[r];for(var a in e)Object.prototype.hasOwnProperty.call(e,a)&&(n[a]=e[a])}return n},W.apply(this,arguments)},Ye=globalThis&&globalThis.__rest||function(n,e){var r={};for(var t in n)Object.prototype.hasOwnProperty.call(n,t)&&e.indexOf(t)<0&&(r[t]=n[t]);if(n!=null&&typeof Object.getOwnPropertySymbols=="function")for(var a=0,t=Object.getOwnPropertySymbols(n);a<t.length;a++)e.indexOf(t[a])<0&&Object.prototype.propertyIsEnumerable.call(n,t[a])&&(r[t[a]]=n[t[a]]);return r},Ze=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},G=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(Ze(arguments[e]));return n};function ke(n,e,r){e===void 0&&(e={}),r===void 0&&(r=[]);var t=e.manual,a=t===void 0?!1:t,i=Ye(e,["manual"]),u=W({manual:a},i),s=J(n),f=Ge(),o=X(function(){var l=r.map(function(c){var g;return(g=c==null?void 0:c.onInit)===null||g===void 0?void 0:g.call(c,u)}).filter(Boolean);return new Qe(s,u,f,Object.assign.apply(Object,G([{}],l)))},[]);return o.options=u,o.pluginImpls=r.map(function(l){return l(o,u)}),Ue(function(){if(!a){var l=o.state.params||e.defaultParams||[];o.run.apply(o,G(l))}}),L(function(){o.cancel()}),{loading:o.state.loading,data:o.state.data,error:o.state.error,params:o.state.params||[],cancel:P(o.cancel.bind(o)),refresh:P(o.refresh.bind(o)),refreshAsync:P(o.refreshAsync.bind(o)),run:P(o.run.bind(o)),runAsync:P(o.runAsync.bind(o)),mutate:P(o.mutate.bind(o))}}var Ie=globalThis&&globalThis.__read||function(n,e){var r=typeof Symbol=="function"&&n[Symbol.iterator];if(!r)return n;var t=r.call(n),a,i=[],u;try{for(;(e===void 0||e-- >0)&&!(a=t.next()).done;)i.push(a.value)}catch(s){u={error:s}}finally{try{a&&!a.done&&(r=t.return)&&r.call(t)}finally{if(u)throw u.error}}return i},er=globalThis&&globalThis.__spread||function(){for(var n=[],e=0;e<arguments.length;e++)n=n.concat(Ie(arguments[e]));return n};function nr(n,e,r){return ke(n,e,er(r||[],[ve,ge,_e,Se,Le,k,se,$e]))}export{nr as u};
