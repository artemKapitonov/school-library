// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {entity} from '../models';

export function CreateBook(arg1:entity.Book):Promise<string>;

export function CreateStudent(arg1:entity.Student):Promise<string>;

export function DeleteBook(arg1:number):Promise<string>;

export function DeleteStudent(arg1:number):Promise<string>;

export function GetAllBooks():Promise<Array<entity.Book>>;

export function GetAllStudents():Promise<Array<entity.Student>>;

export function RegisterStudentBook(arg1:number,arg2:number):Promise<string>;

export function SearchBook(arg1:string):Promise<Array<entity.Book>>;

export function SearchStudent(arg1:string):Promise<Array<entity.Student>>;

export function UnregisterStudentBook(arg1:number,arg2:number):Promise<string>;

export function UpdateBook(arg1:entity.Book):Promise<string>;

export function UpdateStudent(arg1:entity.Student):Promise<string>;
