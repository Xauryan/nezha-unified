export const getBasePath=()=>"/dashboard/certimate/";
export const withBasePath=(path:string)=>/^[a-z][a-z\d+.-]*:/i.test(path)||path.startsWith("//")?path:getBasePath()+path.replace(/^\/+/,"");
