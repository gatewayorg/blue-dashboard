import { useCallback, useState } from 'react';

export function useDialogState(): [boolean, () => void, () => void] {
  const [open, setOpen] = useState(false);

  const openDialog = useCallback(() => {
    setOpen(true);
  }, [setOpen]);

  const closeDialog = useCallback(() => {
    setOpen(false);
  }, [setOpen]);
  return [open, openDialog, closeDialog];
}
