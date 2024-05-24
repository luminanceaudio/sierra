import { ExclamationCircleIcon } from '@heroicons/react/24/solid';
import { StyledToast } from './Toast.style';

export type ToastProps = {
  error: any;
};

function Toast({ error }: ToastProps) {
  if (!error) return null;

  return (
    <StyledToast
      style={{ opacity: 0 }}
      loadStyle={{ opacity: 1 }}
      durationMs={2000}
    >
      <ExclamationCircleIcon width={20} color="#cf1212" />
      {error}
    </StyledToast>
  );
}

export default Toast;
