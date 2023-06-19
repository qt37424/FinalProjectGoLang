export class User {
  constructor(
    public id: string,
    public email: string,
    public password: string,
    public role: string,
    private _token: string,
    private _tokenExpirationDate: Date
  ) {}

  //check if token is expired
  get token() {
    if (!this._tokenExpirationDate || new Date() > this._tokenExpirationDate) {
      return null;
    }
    return this._token;
  }
}
