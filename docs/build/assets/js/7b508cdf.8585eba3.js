"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[960],{3905:(e,t,a)=>{a.d(t,{Zo:()=>c,kt:()=>m});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function i(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function l(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?i(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function o(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},i=Object.keys(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var s=n.createContext({}),p=function(e){var t=n.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):l(l({},t),e)),a},c=function(e){var t=p(e.components);return n.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,i=e.originalType,s=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),d=p(a),m=r,f=d["".concat(s,".").concat(m)]||d[m]||u[m]||i;return a?n.createElement(f,l(l({ref:t},c),{},{components:a})):n.createElement(f,l({ref:t},c))}));function m(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=a.length,l=new Array(i);l[0]=d;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o.mdxType="string"==typeof e?e:r,l[1]=o;for(var p=2;p<i;p++)l[p]=a[p];return n.createElement.apply(null,l)}return n.createElement.apply(null,a)}d.displayName="MDXCreateElement"},5162:(e,t,a)=>{a.d(t,{Z:()=>l});var n=a(7294),r=a(6010);const i="tabItem_Ymn6";function l(e){let{children:t,hidden:a,className:l}=e;return n.createElement("div",{role:"tabpanel",className:(0,r.Z)(i,l),hidden:a},t)}},5488:(e,t,a)=>{a.d(t,{Z:()=>m});var n=a(7462),r=a(7294),i=a(6010),l=a(2389),o=a(7392),s=a(7094),p=a(2466);const c="tabList__CuJ",u="tabItem_LNqP";function d(e){var t;const{lazy:a,block:l,defaultValue:d,values:m,groupId:f,className:g}=e,h=r.Children.map(e.children,(e=>{if((0,r.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})),k=m??h.map((e=>{let{props:{value:t,label:a,attributes:n}}=e;return{value:t,label:a,attributes:n}})),b=(0,o.l)(k,((e,t)=>e.value===t.value));if(b.length>0)throw new Error(`Docusaurus error: Duplicate values "${b.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`);const v=null===d?d:d??(null==(t=h.find((e=>e.props.default)))?void 0:t.props.value)??h[0].props.value;if(null!==v&&!k.some((e=>e.value===v)))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${v}" but none of its children has the corresponding value. Available values are: ${k.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);const{tabGroupChoices:N,setTabGroupChoices:y}=(0,s.U)(),[w,A]=(0,r.useState)(v),T=[],{blockElementScrollPositionUntilNextRender:E}=(0,p.o5)();if(null!=f){const e=N[f];null!=e&&e!==w&&k.some((t=>t.value===e))&&A(e)}const C=e=>{const t=e.currentTarget,a=T.indexOf(t),n=k[a].value;n!==w&&(E(t),A(n),null!=f&&y(f,String(n)))},I=e=>{var t;let a=null;switch(e.key){case"Enter":C(e);break;case"ArrowRight":{const t=T.indexOf(e.currentTarget)+1;a=T[t]??T[0];break}case"ArrowLeft":{const t=T.indexOf(e.currentTarget)-1;a=T[t]??T[T.length-1];break}}null==(t=a)||t.focus()};return r.createElement("div",{className:(0,i.Z)("tabs-container",c)},r.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,i.Z)("tabs",{"tabs--block":l},g)},k.map((e=>{let{value:t,label:a,attributes:l}=e;return r.createElement("li",(0,n.Z)({role:"tab",tabIndex:w===t?0:-1,"aria-selected":w===t,key:t,ref:e=>T.push(e),onKeyDown:I,onClick:C},l,{className:(0,i.Z)("tabs__item",u,null==l?void 0:l.className,{"tabs__item--active":w===t})}),a??t)}))),a?(0,r.cloneElement)(h.filter((e=>e.props.value===w))[0],{className:"margin-top--md"}):r.createElement("div",{className:"margin-top--md"},h.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==w})))))}function m(e){const t=(0,l.Z)();return r.createElement(d,(0,n.Z)({key:String(t)},e))}},5364:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>c,contentTitle:()=>s,default:()=>m,frontMatter:()=>o,metadata:()=>p,toc:()=>u});var n=a(7462),r=(a(7294),a(3905)),i=(a(8209),a(5488)),l=a(5162);const o={description:"Connect to a Microsft 365 tenant"},s="Microsoft 365 access",p={unversionedId:"setup/m365_access",id:"setup/m365_access",title:"Microsoft 365 access",description:"Connect to a Microsft 365 tenant",source:"@site/docs/setup/m365_access.md",sourceDirName:"setup",slug:"/setup/m365_access",permalink:"/docs/setup/m365_access",draft:!1,editUrl:"https://github.com/alcionai/corso/tree/main/docs/docs/setup/m365_access.md",tags:[],version:"current",frontMatter:{description:"Connect to a Microsft 365 tenant"}},c={},u=[{value:"Create an Azure AD application",id:"create-an-azure-ad-application",level:2},{value:"Register a new application",id:"register-a-new-application",level:3},{value:"Configure basic settings",id:"configure-basic-settings",level:3},{value:"Configure required permissions",id:"configure-required-permissions",level:3},{value:"Grant admin consent",id:"grant-admin-consent",level:3},{value:"Export application credentials",id:"export-application-credentials",level:2},{value:"Tenant ID and client ID",id:"tenant-id-and-client-id",level:3},{value:"Azure client secret",id:"azure-client-secret",level:3}],d={toc:u};function m(e){let{components:t,...o}=e;return(0,r.kt)("wrapper",(0,n.Z)({},d,o,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"microsoft-365-access"},"Microsoft 365 access"),(0,r.kt)("p",null,"To perform backup and restore operations, Corso requires access to your ",(0,r.kt)("a",{parentName:"p",href:"concepts#m365-concepts"},"M365 tenant"),"\nby creating an ",(0,r.kt)("a",{parentName:"p",href:"concepts#m365-concepts"},"Azure AD application")," with appropriate permissions."),(0,r.kt)("p",null,"The following steps outline a simplified procedure for creating an Azure Ad application suitable for use with Corso.\nFor more details, please refer to the\n",(0,r.kt)("a",{parentName:"p",href:"https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal"},"official documentation"),"\nfor adding an Azure AD Application and Service Principal using the Azure Portal."),(0,r.kt)("h2",{id:"create-an-azure-ad-application"},"Create an Azure AD application"),(0,r.kt)("p",null,"Sign in into the ",(0,r.kt)("a",{parentName:"p",href:"https://portal.azure.com/"},"Azure Portal")," with a user that has sufficient permissions to create an\nAD application."),(0,r.kt)("h3",{id:"register-a-new-application"},"Register a new application"),(0,r.kt)("p",null,"From the list of ",(0,r.kt)("a",{parentName:"p",href:"https://portal.azure.com/#allservices"},"Azure services"),", select\n",(0,r.kt)("strong",{parentName:"p"},"Azure Active Directory ","\u2192"," App Registrations ","\u2192"," New Registration")),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Registering a new application",src:a(1256).Z,width:"1760",height:"1322"})),(0,r.kt)("h3",{id:"configure-basic-settings"},"Configure basic settings"),(0,r.kt)("p",null,"Next, configure the following:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Give the application a name"),(0,r.kt)("li",{parentName:"ul"},"Select ",(0,r.kt)("strong",{parentName:"li"},"Accounts in this organizational directory only")),(0,r.kt)("li",{parentName:"ul"},"Skip the ",(0,r.kt)("strong",{parentName:"li"},"Redirect URI")," option"),(0,r.kt)("li",{parentName:"ul"},"Click ",(0,r.kt)("strong",{parentName:"li"},"Register")," at the bottom of the screen")),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Configuring the application",src:a(3855).Z,width:"1820",height:"1400"})),(0,r.kt)("h3",{id:"configure-required-permissions"},"Configure required permissions"),(0,r.kt)("p",null,"Within the new application (",(0,r.kt)("inlineCode",{parentName:"p"},"CorsoApp")," in the below diagram), select ",(0,r.kt)("strong",{parentName:"p"},"API Permissions ","\u2192"," Add a permission")," from\nthe management panel."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Adding application permissions",src:a(8786).Z,width:"1940",height:"1266"})),(0,r.kt)("p",null,"Select the following permissions from ",(0,r.kt)("strong",{parentName:"p"},"Microsoft API ","\u2192"," Microsoft Graph ","\u2192"," Application Permissions")," and\nthen click ",(0,r.kt)("strong",{parentName:"p"},"Add permissions"),"."),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:"left"},"API / Permissions Name"),(0,r.kt)("th",{parentName:"tr",align:"left"},"Type"),(0,r.kt)("th",{parentName:"tr",align:"left"},"Description"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:"left"},"Calendars.ReadWrite"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Application"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Read and write calendars in all mailboxes")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:"left"},"Contacts.ReadWrite"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Application"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Read and write contacts in all mailboxes")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:"left"},"Files.ReadWrite.All"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Application"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Read and write files in all site collections")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:"left"},"Mail.ReadWrite"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Application"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Read and write mail in all mailboxes")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:"left"},"User.Read.All"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Application"),(0,r.kt)("td",{parentName:"tr",align:"left"},"Read all users' full profiles")))),(0,r.kt)("h3",{id:"grant-admin-consent"},"Grant admin consent"),(0,r.kt)("p",null,"Finally, grant admin consent to this application. This step is required even if the user that created the application\nis an Microsoft 365 admin."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Granting administrator consent",src:a(5853).Z,width:"1940",height:"1266"})),(0,r.kt)("h2",{id:"export-application-credentials"},"Export application credentials"),(0,r.kt)("p",null,"After configuring the Corso Azure AD application, store the information needed by Corso to connect to the application\nas environment variables."),(0,r.kt)("h3",{id:"tenant-id-and-client-id"},"Tenant ID and client ID"),(0,r.kt)("p",null,"To view the tenant and client ID, select Overview from the app management panel."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Obtaining Tenant and Client IDs",src:a(7586).Z,width:"1960",height:"1132"})),(0,r.kt)("p",null,"Copy the client and tenant IDs and export them into the following environment variables."),(0,r.kt)(i.Z,{groupId:"os",mdxType:"Tabs"},(0,r.kt)(l.Z,{value:"win",label:"Powershell",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-powershell"},'$Env:AZURE_CLIENT_ID = "<Application (client) ID for configured app>"\n$Env:AZURE_TENANT_ID = "<Directory (tenant) ID for configured app>"\n'))),(0,r.kt)(l.Z,{value:"unix",label:"Linux/macOS",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>\nexport AZURE_CLIENT_ID=<Application (client) ID for configured app>\n"))),(0,r.kt)(l.Z,{value:"docker",label:"Docker",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"export AZURE_TENANT_ID=<Directory (tenant) ID for configured app>\nexport AZURE_CLIENT_ID=<Application (client) ID for configured app>\n")))),(0,r.kt)("h3",{id:"azure-client-secret"},"Azure client secret"),(0,r.kt)("p",null,"Finally, you need to obtain a client secret associated with the app using ",(0,r.kt)("strong",{parentName:"p"},"Certificates & Secrets")," from the app\nmanagement panel."),(0,r.kt)("p",null,"Click ",(0,r.kt)("strong",{parentName:"p"},"New Client Secret")," under ",(0,r.kt)("strong",{parentName:"p"},"Client secrets")," and follow the instructions to create a secret."),(0,r.kt)("p",null,(0,r.kt)("img",{alt:"Obtaining the Azure client secrete",src:a(2).Z,width:"2714",height:"1502"})),(0,r.kt)("p",null,"After creating the secret, immediately copy the secret ",(0,r.kt)("strong",{parentName:"p"},"Value")," because it won't be available later. Export it as an\nenvironment variable."),(0,r.kt)(i.Z,{groupId:"os",mdxType:"Tabs"},(0,r.kt)(l.Z,{value:"win",label:"Powershell",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-powershell"},'$Env:AZURE_CLIENT_SECRET = "<Client secret value>"\n'))),(0,r.kt)(l.Z,{value:"unix",label:"Linux/macOS",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"export AZURE_CLIENT_SECRET=<Client secret value>\n"))),(0,r.kt)(l.Z,{value:"docker",label:"Docker",mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"export AZURE_CLIENT_SECRET=<Client secret value>\n")))))}m.isMDXComponent=!0},3855:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_configure-83e59417c5e7b0ed87fa616327edb2fc.png"},5853:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_consent-a3314620736ed126f05fbe7baebbb320.png"},1256:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_create_new-097b53ba552b60a675f563cc7b60d904.png"},7586:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_ids-ad4bede65b17113d84906a8d8723464f.png"},8786:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_permissions-b64c40b5f25e4a3ce772065e2d39e774.png"},2:(e,t,a)=>{a.d(t,{Z:()=>n});const n=a.p+"assets/images/m365app_secret-fc21706893de97d86d9efc50fc248987.png"},8209:(e,t,a)=>{a(7294)}}]);