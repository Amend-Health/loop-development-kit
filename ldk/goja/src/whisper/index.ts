export enum WhisperComponentType {
  BOX = 'box',
  BUTTON = 'button',
  CHECKBOX = 'checkbox',
  COLLAPSE_BOX = 'collapseBox',
  DIVIDER = 'divider',
  LINK = 'link',
  LIST_PAIR = 'listPair',
  MARKDOWN = 'markdown',
  MESSAGE = 'message',
  NUMBER = 'number',
  PASSWORD = 'password',
  RADIO_GROUP = 'radioGroup',
  SELECT = 'select',
  TELEPHONE = 'telephone',
  TEXT_INPUT = 'textInput',
}

export enum Alignment {
  CENTER = 'center',
  LEFT = 'left',
  RIGHT = 'right',
  SPACE_AROUND = 'space_around',
  SPACE_EVENLY = 'space_evenly',
}

export enum Direction {
  HORIZONTAL = 'horizontal',
  VERTICAL = 'vertical',
}

export enum TextAlign {
  CENTER = 'center',
  LEFT = 'left',
  RIGHT = 'right',
}

export enum Urgency {
  ERROR = 'error',
  NONE = 'none',
  SUCCESS = 'success',
  WARNING = 'warning',
}

export interface WhisperComponent {
  id: string;
  type: string;
}

export interface Button extends WhisperComponent {
  label: string;
  onClick(): Promise<void>;
  submit?: boolean;
}

export interface Checkbox extends WhisperComponent {
  label: string;
  tooltip?: string;
  value?: boolean;
}

export interface Link extends WhisperComponent {
  href?: string;
  style?: string;
  onClick?(): Promise<string>;
  text: string;
  textAlign?: string;
}

export interface ListPair extends WhisperComponent {
  copyable: boolean;
  label: string;
  style?: string;
  value?: string;
}

export interface Markdown extends WhisperComponent {
  body: string;
}

export interface Message extends WhisperComponent {
  body: string;
  header: string;
  style?: Urgency;
  textAlign?: TextAlign;
}

export interface NumberInput extends WhisperComponent {
  label: string;
  max?: number;
  min?: number;
  onChange?(): Promise<number>;
  step?: number;
  tooltip?: string;
  value?: number;
}

export interface Password extends WhisperComponent {
  label: string;
  onChange?(): Promise<string>;
  tooltip?: string;
  value?: string;
}

export interface RadioGroup extends WhisperComponent {
  onSelect(): Promise<any>;
  options: Record<string, any>;
  selected?: any;
}

export interface Select extends WhisperComponent {
  label: string;
  onChange?(): Promise<any>;
  tooltip?: string;
  options: Record<string, any>;
}

export interface Telephone extends WhisperComponent {
  label: string;
  onChange?(): Promise<string>;
  pattern?: RegExp;
  tooltip?: string;
  value?: string;
}

export interface TextInput extends WhisperComponent {
  label: string;
  onChange?(): Promise<string>;
  tooltip?: string;
  value?: string;
}

export interface CollapseBox extends WhisperComponent {
  children: Array<
    | Button
    | Checkbox
    | Link
    | ListPair
    | Markdown
    | Message
    | NumberInput
    | Password
    | RadioGroup
    | Select
    | Telephone
    | TextInput
  >;
  label?: string;
  open: boolean;
}

export interface Box extends WhisperComponent {
  alignment: string;
  children: Array<
    | Button
    | Checkbox
    | CollapseBox
    | Link
    | ListPair
    | Markdown
    | Message
    | NumberInput
    | Password
    | RadioGroup
    | Select
    | Telephone
    | TextInput
  >;
  direction: string;
}

/**
 * The HTTP Request configuration.
 */
export interface NewWhisper {
  components: Array<
    | Box
    | Button
    | Checkbox
    | CollapseBox
    | Link
    | ListPair
    | Markdown
    | Message
    | NumberInput
    | Password
    | RadioGroup
    | Select
    | Telephone
    | TextInput
  >;
  label: string;
}

export interface Whisper {
  components: Array<
    | Box
    | Button
    | Checkbox
    | CollapseBox
    | Link
    | ListPair
    | Markdown
    | Message
    | NumberInput
    | Password
    | RadioGroup
    | Select
    | Telephone
    | TextInput
  >;
  close(callback: () => void): void;
  id: string;
  label: string;
  onChange(callback: (whisper: Whisper) => void): void;
}

export interface WhisperAptitude {
  all(): Promise<Whisper[]>;
  create(whisper: NewWhisper, cb: (whisper: Whisper) => void): void;
}

function all(): Promise<Whisper[]> {
  return new Promise<Whisper[]>((resolve, reject) => {
    try {
      oliveHelps.whisper.all((whispers: Whisper[]) => resolve(whispers));
    } catch (error) {
      console.log(error);
      reject(error);
    }
  });
}

function create(whisper: NewWhisper, cb: (whisper: Whisper) => void): void {
  return oliveHelps.whisper.create(whisper, cb);
}

export const whisper: WhisperAptitude = {
  all,
  create,
};
