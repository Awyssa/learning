import { useEffect, useRef } from "react";
import { createPortal } from "react-dom";

const Modal = ({ children }) => {

  const elRef = useRef(null);

  if(!elRef.current) { 
    // Maintain the same div per each render
    elRef.current = document.createElement("div")
  }

  useEffect(() => {
    const modalRoot = document.getElementById("modal")
    modalRoot?.appendChild(elRef.current)

    return true
  }, [])

  return createPortal(<div>{children}</div>, elRef.current)
}

export default Modal;