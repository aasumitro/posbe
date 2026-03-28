import {z} from "zod";

export const FormSchema = z.object({
  name: z.string().trim().min(1, "Name is required"),
  phone: z.string()
    .regex(/^\+[1-9]\d{6,14}$/, "Invalid phone number format. Please use international format, e.g: +628123456789")
    .or(z.literal("")),
  email: z.string(),
  description: z.string(),
}).superRefine((data, ctx) => {
  if (data.email !== "" && !z.string().email().safeParse(data.email).success) {
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: "Invalid email address",
      path: ["email"],
    });
  }

  if (!data.phone.trim() && !data.email.trim()) {
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: "Phone or email is required",
      path: ["phone"],
    });
    ctx.addIssue({
      code: z.ZodIssueCode.custom,
      message: "Phone or email is required",
      path: ["email"],
    });
  }
});

export type FormValues = z.infer<typeof FormSchema>;

export type CustomerErrorResponse = {
  name?: string[]
  phone?: string[]
  email?: string[]
  description?: string[]
}