import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

interface AuthRequest extends Request {
    user?: any;
}

export const authorization = (req: AuthRequest, res: Response, next: NextFunction): Response | void => {
  let token: string = "";
  if (req.headers.authorization && req.headers.authorization.startsWith("Bearer ")) {
    let array: string[] = req.headers.authorization.split(" ");
    token  = array[1];
  } else {
    return res.status(401).send("Access denied.");
  }

  try {
    const decoded: any = jwt.verify(token, "axdaasdFWFQWF232SegqwQrqww");
    req.user = decoded;
    next();
  } catch (ex) {
    console.log(ex)
    return res.status(401).send("Access deniedd.");
  }
};

