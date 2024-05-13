import React, { useEffect, useState } from 'react';
import { StyledTransitionOnLoad } from './TransitionOnLoad.style';

export type TransitionOnLoadProps = {
  children: React.ReactNode;
  style: React.CSSProperties;
  loadStyle: React.CSSProperties;
  className?: string;
  durationMs?: number;
};

function TransitionOnLoad({
  children,
  style,
  loadStyle,
  className,
  durationMs = 0,
}: TransitionOnLoadProps) {
  const [hasLoaded, setHasLoaded] = React.useState(false);
  const [timeLeft, setTimeLeft] = useState<number>(-1);

  // Load
  useEffect(() => {
    const intervalId = setInterval(() => {
      setHasLoaded(true);
    }, 1);

    // Cleanup
    return () => clearInterval(intervalId);
  }, []);

  // Duration timer
  useEffect(() => {
    setTimeLeft(durationMs);

    const intervalId = setInterval(() => {
      setTimeLeft(0);
    }, durationMs);

    // Cleanup
    return () => clearInterval(intervalId);
  }, [durationMs]);

  return (
    <StyledTransitionOnLoad
      style={{ ...style, ...(hasLoaded && timeLeft !== 0 ? loadStyle : {}) }}
      className={className}
    >
      {children}
    </StyledTransitionOnLoad>
  );
}

TransitionOnLoad.defaultProps = {
  className: undefined,
  durationMs: 0,
};

export default TransitionOnLoad;
