import React, {
  KeyboardEventHandler,
  forwardRef,
  useEffect,
  useRef,
} from 'react';

export type TextareaProps = {
  value: string;
  onChange: (value: string) => void;
  onBlur?: () => void;
  onFocus?: () => void;
  rows?: number;
  placeholder?: string;
  style?: React.CSSProperties;
  onKeyDown?: KeyboardEventHandler<HTMLTextAreaElement>;
  getRef: () => any;
  className?: string;
};

const Textarea = forwardRef(function Textarea(
  {
    value,
    onChange,
    rows,
    placeholder,
    style,
    onKeyDown,
    onBlur,
    getRef,
    onFocus,
    className,
  }: TextareaProps,
  ref: any,
) {
  const usedRef = useRef();
  const textareaRef: any = ref || usedRef;

  let finalRows: number | undefined = rows || value.split('\n').length;
  if (finalRows === 0) {
    finalRows = 1;
  }

  useEffect(() => {
    const refX = getRef();
    refX.style.height = '5px';
    refX.style.height = `${refX.scrollHeight - 24}px`;
  }, []);

  return (
    <textarea
      ref={textareaRef}
      style={{
        background: 'none',
        color: '#FFF',
        flex: 1,
        fontSize: 14,
        width: '100%',
        overflow: 'hidden',
        ...style,
      }}
      rows={1}
      value={value}
      placeholder={placeholder}
      onChange={(e) => {
        onChange(e.target.value);

        // Scroll all the way down
        // e.currentTarget.scrollTop = e.currentTarget.scrollHeight;

        // Auto-grow: Reset the height, get new scroll height and set it as the height
        e.currentTarget.style.height = '5px';
        e.currentTarget.style.height = `${e.currentTarget.scrollHeight - 24}px`;
      }}
      onKeyDown={onKeyDown}
      onBlur={onBlur}
      className={className}
      autoComplete="off"
      autoCorrect="off"
      autoCapitalize="off"
      spellCheck={false}
      onFocus={onFocus}
    />
  );
});

Textarea.defaultProps = {
  rows: 0,
  placeholder: '',
  style: {},
  onKeyDown: () => {},
  onBlur: () => {},
  onFocus: () => {},
  className: undefined,
};

export default Textarea;
