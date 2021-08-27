import {
  Box,
  BoxChildComponent,
  Component,
  NewWhisper,
  StateMap,
  UpdateWhisper,
  Whisper,
  WhisperComponentType,
  File,
} from './types';

function throwForDuplicateKeys<T extends { key?: string }>(components: T[]): T[] {
  const keySet = new Set<string>();
  components.forEach((item) => {
    if (item.key === undefined) {
      return;
    }
    if (keySet.has(item.key)) {
      throw new Error(
        `Duplicate Key ${item.key} encountered, provided keys must be unique among its siblings`,
      );
    }
    keySet.add(item.key);
  });
  return components;
}

export function mapToInternalChildComponent(
  component: BoxChildComponent,
  stateMap: StateMap,
): OliveHelps.ChildComponents {
  const onClick = 'onClick' in component ? component.onClick : null;
  switch (component.type) {
    case WhisperComponentType.Box:
      // eslint-disable-next-line no-case-declarations
      const { justifyContent, ...otherProps } = component as Box;
      return {
        ...otherProps,
        alignment: 'justifyContent' in component ? justifyContent : component.alignment,
        children: throwForDuplicateKeys(
          component.children.map((childComponent) =>
            mapToInternalChildComponent(childComponent, stateMap),
          ),
        ),
        type: WhisperComponentType.Box,
        onClick: onClick
          ? (error, whisper) => onClick(error, mapToExternalWhisper(whisper, stateMap))
          : undefined,
      } as OliveHelps.Box;
    case WhisperComponentType.Button:
      return {
        ...component,
        onClick: (error, whisper) => {
          component.onClick(error, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Button;
    case WhisperComponentType.Checkbox:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Checkbox;
    case WhisperComponentType.Email:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Email;
    case WhisperComponentType.Link: {
      return {
        ...component,
        onClick: onClick
          ? (error, whisper) => onClick(error, mapToExternalWhisper(whisper, stateMap))
          : undefined,
      } as OliveHelps.Link;
    }
    case WhisperComponentType.Divider:
    case WhisperComponentType.SectionTitle:
      return component;
    case WhisperComponentType.RichTextEditor:
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.RichTextEditor;
    case WhisperComponentType.DropZone:
      return {
        ...component,
        onDrop: (error, param, whisper) => {
          const callbackHandler: (file: OliveHelps.File) => File = (file: OliveHelps.File) => ({
            path: file.path,
            size: file.size,
            readFile: () =>
              new Promise<Uint8Array>((resolve, reject) => {
                file.readFile((readError, buffer) => {
                  if (readError) {
                    return reject(readError);
                  }
                  return resolve(new Uint8Array(buffer));
                });
              }),
          });
          component.onDrop(
            error,
            param.map(callbackHandler),
            mapToExternalWhisper(whisper, stateMap),
          );
        },
      };
    case WhisperComponentType.ListPair: {
      // eslint-disable-next-line
      const { onCopy } = component;
      if (onCopy) {
        return {
          ...component,
          onCopy: (error, param, whisper) => {
            onCopy(error, param, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.ListPair;
      }
      return component as OliveHelps.ListPair;
    }
    case WhisperComponentType.Message: {
      // eslint-disable-next-line
      const { onCopy } = component;
      if (onCopy) {
        return {
          ...component,
          onCopy: (error, whisper) => {
            onCopy(error, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.Message;
      }
      return component as OliveHelps.Message;
    }
    case WhisperComponentType.Markdown: {
      // eslint-disable-next-line
      const { onCopy } = component;
      if (onCopy) {
        return {
          ...component,
          onCopy: (error, whisper) => {
            onCopy(error, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.Markdown;
      }
      return component as OliveHelps.Markdown;
    }
    case WhisperComponentType.Number:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.NumberInput;
    case WhisperComponentType.Password:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Password;
    case WhisperComponentType.RadioGroup:
      if (component.id && component.selected) {
        stateMap.set(component.id, component.selected);
      }
      return {
        ...component,
        onSelect: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onSelect(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.RadioGroup;
    case WhisperComponentType.Select:
      if (component.id && component.selected) {
        stateMap.set(component.id, component.selected);
      }
      return {
        ...component,
        onSelect: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onSelect(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Select;
    case WhisperComponentType.Telephone:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.Telephone;
    case WhisperComponentType.TextInput:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value);
      }
      return {
        ...component,
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.TextInput;
    case WhisperComponentType.DateTimeInput:
      if (component.id && component.value) {
        stateMap.set(component.id, component.value.toISOString());
      }
      return {
        ...component,
        value: component.value?.toISOString(),
        max: component.max?.toISOString(),
        min: component.min?.toISOString(),
        onChange: (error, param, whisper) => {
          if (component.id) {
            stateMap.set(component.id, param);
          }
          component.onChange(error, param, mapToExternalWhisper(whisper, stateMap));
        },
      } as OliveHelps.DateTimeInput;
    case WhisperComponentType.Icon:
      // eslint-disable-next-line
      const onIconClick = component.onClick;
      if (onIconClick) {
        return {
          ...component,
          onClick: (error, whisper) => {
            onIconClick(error, mapToExternalWhisper(whisper, stateMap));
          },
        } as OliveHelps.Icon;
      }
      return component as OliveHelps.Icon;
    default:
      // Suppressing warning to deal with unexpected types.
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      throw new Error(`Unexpected component type: ${(component as any)?.type}`);
  }
}

export function mapToInternalComponent(
  component: Component,
  stateMap: StateMap,
): OliveHelps.Components {
  switch (component.type) {
    case WhisperComponentType.CollapseBox:
      // eslint-disable-next-line no-case-declarations
      const onClick = 'onClick' in component ? component.onClick : null;
      return {
        ...component,
        children: throwForDuplicateKeys(
          component.children.map((childComponent) =>
            mapToInternalChildComponent(childComponent, stateMap),
          ),
        ),
        type: WhisperComponentType.CollapseBox,
        onClick: onClick
          ? (error, param, whisper) =>
              onClick(error, param, mapToExternalWhisper(whisper, stateMap))
          : undefined,
      } as OliveHelps.CollapseBox;
    default:
      return mapToInternalChildComponent(component, stateMap);
  }
}

export function mapToInternalWhisper(
  whisper: NewWhisper,
  stateMap: StateMap,
): OliveHelps.NewWhisper;
export function mapToInternalWhisper(
  whisper: UpdateWhisper,
  stateMap: StateMap,
): OliveHelps.UpdateWhisper;
export function mapToInternalWhisper(
  whisper: NewWhisper | UpdateWhisper,
  stateMap: StateMap,
): OliveHelps.NewWhisper | OliveHelps.UpdateWhisper {
  return 'onClose' in whisper
    ? {
        label: whisper.label,
        onClose: whisper.onClose,
        components: throwForDuplicateKeys(
          whisper.components.map((component) => mapToInternalComponent(component, stateMap)),
        ),
      }
    : {
        label: whisper.label,
        components: throwForDuplicateKeys(
          whisper.components.map((component) => mapToInternalComponent(component, stateMap)),
        ),
      };
}

export function mapToExternalWhisper(whisper: OliveHelps.Whisper, stateMap: StateMap): Whisper {
  return {
    id: whisper.id,
    close: whisper.close,
    componentState: stateMap,
    update(updateWhisper: UpdateWhisper, cb): void {
      whisper.update(mapToInternalWhisper(updateWhisper, stateMap), cb);
    },
  };
}
