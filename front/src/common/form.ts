import { FieldRenderProps } from "react-final-form";

export function hasError(meta: FieldRenderProps<HTMLElement>["meta"]) {
  return (
    ((meta.submitError && !meta.dirtySinceLastSubmit) || meta.error) &&
    meta.touched
  );
}

export function getErrorText(meta: FieldRenderProps<HTMLElement>["meta"]) {
  return !!hasError(meta) ? meta.error || meta.submitError : undefined;
}
