import {Button} from "@/components/ui/button";
import {IconChevronLeft, IconPlus, IconTrash} from "@tabler/icons-react";
import {useNavigate} from "@tanstack/react-router";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Separator} from "@/components/ui/separator";
import ReactQuill from 'react-quill-new';
import 'react-quill-new/dist/quill.snow.css';
import {useState} from "react";

export function NewProductPage() {
  const navigate = useNavigate();
  const [value, setValue] = useState('');

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-4">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="flex gap-3 items-center">
          <Button
            variant="outline"
            className="w-8 h-8 cursor-pointer"
            onClick={async (e) => {
              e.preventDefault();
              await navigate({to: "/stores/catalogs"})
            }}
          >
            <IconChevronLeft className="w-4 h-4" />
          </Button>
          <h5 className="text-lg font-semibold">
            Add New Product
          </h5>
        </div>
        <div className="flex gap-2 items-center">
          <Button variant="ghost" className="cursor-pointer">Save as Draft</Button>
          <Button className="cursor-pointer">Publish</Button>
        </div>
      </aside>

      <aside className="mx-4 gap-6">
        <div className="grid grid-cols-2 gap-12 p-8">
          <div className="w-full space-y-4 mb-10">
            <h4 className="text-md font-semibold">Product Image</h4>

            <div className="flex items-start gap-4">
              {/* Upload Area */}
              <label
                htmlFor="image-upload"
                className="w-52 h-52 border-2 border-dashed rounded-lg cursor-pointer flex flex-col items-center justify-center text-gray-500 hover:border-black"
              >
                <svg className="w-6 h-6 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                </svg>
                <p className="text-sm text-center">
                  Drop your image here<br />
                  or <span className="text-blue-600 underline">Click to browse</span>
                </p>
                <input
                  type="file"
                  id="image-upload"
                  accept="image/jpeg,image/png,image/jpg"
                  // onChange={handleImageChange}
                  className="hidden"
                />
              </label>
            </div>

            <div className="flex text-xs text-gray-500">
              <span className="inline-block mr-1">ℹ️</span>
              <p>Only image of .jpeg, .jpg, .png and a minimum size of 250 x 250px for optimal image use.</p>
            </div>

            <Separator className="my-6" />

            <h4 className="text-md font-semibold">Product Information</h4>

            <div className="grid w-full items-center gap-3">
              <Label htmlFor="code" className="text-muted-foreground">Code</Label>
              <Input
                type="text"
                id="code"
                placeholder="Enter the (UPC, SKU, etc) Code"
                className="h-12"
              />
            </div>

            <div className="grid w-full items-center gap-3">
              <Label htmlFor="code" className="text-muted-foreground">Name</Label>
              <Input
                type="text"
                id="name"
                placeholder="Enter the Name"
                className="h-12"
              />
            </div>

            <div className="flex gap-2">
              <div className="grid w-full items-center gap-3">
                <Label htmlFor="code" className="text-muted-foreground">Category</Label>
                <Input
                  type="text"
                  id="category"
                  placeholder="Enter the Category"
                  className="h-12"
                />
              </div>
              <div className="grid w-full items-center gap-3">
                <Label htmlFor="code" className="text-muted-foreground">Subcategory</Label>
                <Input
                  type="text"
                  id="subcategory"
                  placeholder="Enter the Subcategory"
                  className="h-12"
                />
              </div>
            </div>

            <div className="grid w-full items-center gap-3">
              <Label htmlFor="code" className="text-muted-foreground">Descriptions</Label>
              <ReactQuill
                theme="snow"
                value={value}
                onChange={setValue}
                className="h-64"
              />
            </div>
          </div>

          <div className="w-full space-y-4">
            <h4 className="text-md font-semibold">Product Variant</h4>

            {[1,2].map((item) => (
              <div key={item} className="mt-4 gap-6 border rounded-xl">
                <div className="bg-gray-50 p-4 rounded-t-xl flex items-center justify-between">
                  <h4 className="text-md">Variant {item}</h4>

                  {item === 2 && (
                    <Button variant="ghost" className="text-red-500 hover:text-red-600 cursor-pointer">
                      <IconTrash  className="w-4 h-4" />
                    </Button>
                  )}
                </div>

                <div className="px-4 py-6 space-y-4">
                  <div className="flex gap-2">
                    <div className="grid w-full items-center gap-3">
                      <Label htmlFor="code" className="text-muted-foreground">Type</Label>
                      <Input
                        type="text"
                        id="category"
                        placeholder="Enter the Type"
                        className="h-12"
                      />
                    </div>
                    <div className="grid w-full items-center gap-3">
                      <Label htmlFor="code" className="text-muted-foreground">Name</Label>
                      <Input
                        type="text"
                        id="subcategory"
                        placeholder="Enter the Name"
                        className="h-12"
                      />
                    </div>
                  </div>

                  <div className="grid w-full items-center gap-3">
                    <Label htmlFor="code" className="text-muted-foreground">Description</Label>
                    <Input
                      type="text"
                      id="subcategory"
                      placeholder="Enter the Description"
                      className="h-12"
                    />
                  </div>

                  <div className="flex gap-2">
                    <div className="grid w-full items-center gap-3">
                      <Label htmlFor="code" className="text-muted-foreground">Unit</Label>
                      <Input
                        type="text"
                        id="category"
                        placeholder="Enter the Unit"
                        className="h-12"
                      />
                    </div>
                    <div className="grid w-full items-center gap-3">
                      <Label htmlFor="code" className="text-muted-foreground">Size</Label>
                      <Input
                        type="text"
                        id="subcategory"
                        placeholder="Enter the Unit Size"
                        className="h-12"
                      />
                    </div>
                  </div>

                  <div className="grid w-full items-center gap-3">
                    <Label htmlFor="code" className="text-muted-foreground">Price</Label>
                    <Input
                      type="text"
                      id="subcategory"
                      placeholder="Enter the Description"
                      className="h-12"
                    />
                  </div>
                </div>
              </div>
            ))}


            <div className="flex justify-center">
              <Button variant="ghost" className="cursor-pointer">
                <IconPlus className="w-4 h-4" />
                Add variant
              </Button>
            </div>
          </div>
        </div>
      </aside>
    </div>
  )
}