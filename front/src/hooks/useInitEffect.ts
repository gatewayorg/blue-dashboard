import React, { useEffect } from 'react';
function useInitEffect(effect: React.EffectCallback): void {
  // eslint-disable-next-line
  return useEffect(effect, []);
}

export { useInitEffect };