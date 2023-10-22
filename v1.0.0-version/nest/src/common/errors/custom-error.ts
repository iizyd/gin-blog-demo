export class CustomError {
  private readonly code: number;
  private readonly msg: string;

  constructor(code: number, msg: string) {
    this.code = code;
    this.msg = msg;
  }

  getCode() {
    return this.code;
  }

  getMsg() {
    return this.msg;
  }
}

export function createError(code: number, msg: string): CustomError {
  return new CustomError(code, msg);
}
