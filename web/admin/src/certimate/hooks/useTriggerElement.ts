import { cloneElement, createElement, isValidElement, useMemo } from "react";
export type UseTriggerElementOptions={onClick?:(event:React.MouseEvent)=>void};
const useTriggerElement=(trigger:React.ReactNode,options?:UseTriggerElementOptions)=>{
 const onClick=options?.onClick;
 return useMemo(()=>{
  if(!trigger)return null;
  if(!isValidElement<UseTriggerElementOptions>(trigger))return createElement("span",{onClick},trigger);
  return cloneElement(trigger,{...trigger.props,onClick:(event:React.MouseEvent)=>{onClick?.(event);trigger.props.onClick?.(event);}});
 },[trigger,onClick]);
};
export default useTriggerElement;
