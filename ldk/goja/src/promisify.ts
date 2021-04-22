import { Cancellable } from './cancellable';

export function promisify<T>(arg: OliveHelps.Readable<T>): Promise<T> {
  return new Promise((resolve, reject) => {
    try {
      arg((cb: T) => resolve(cb));
    } catch (e) {
      console.error(e);
      reject(e);
    }
  });
}

export function promisifyWithParam<TParam, TOut>(
  param: TParam,
  arg: OliveHelps.ReadableWithParam<TParam, TOut>,
): Promise<TOut> {
  return new Promise((resolve, reject) => {
    try {
      arg(param, (cb: TOut) => resolve(cb));
    } catch (e) {
      console.error(e);
      reject(e);
    }
  });
}

export function promisifyWithTwoParams<TParam1, TParam2, TOut>(
  param: TParam1,
  param2: TParam2,
  arg: OliveHelps.ReadableWithTwoParams<TParam1, TParam2, TOut>,
): Promise<TOut> {
  return new Promise((resolve, reject) => {
    try {
      arg(param, param2, (cb: TOut) => resolve(cb));
    } catch (e) {
      console.error(e);
      reject(e);
    }
  });
}

export function promisifyWithFourParams<TP1, TP2, TP3, TP4, TOut>(
  p1: TP1,
  p2: TP2,
  p3: TP3,
  p4: TP4,
  arg: OliveHelps.ReadableWithFourParams<TP1, TP2, TP3, TP4, TOut>,
): Promise<TOut> {
  return new Promise((resolve, reject) => {
    try {
      arg(p1, p2, p3, p4, (cb: TOut) => resolve(cb));
    } catch (e) {
      console.error(e);
      reject(e);
    }
  });
}

export function promisifyListenable<T>(
  cb: (v: T) => void,
  arg: OliveHelps.Listenable<T>,
): Promise<Cancellable> {
  return new Promise((resolve, reject) => {
    try {
      arg(cb, (obj) => {
        resolve(obj);
      });
    } catch (e) {
      console.error('Received error making request', e);
      reject(e);
    }
  });
}

export function promisifyListenableWithParam<TParam, TOut>(
  param: TParam,
  cb: (v: TOut) => void,
  arg: OliveHelps.ListenableWithParam<TParam, TOut>,
): Promise<Cancellable> {
  return new Promise((resolve, reject) => {
    try {
      arg(param, cb, (obj) => {
        resolve(obj);
      });
    } catch (e) {
      console.error('Received error making request', e);
      reject(e);
    }
  });
}
