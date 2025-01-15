# Configuration Instructions

## Setup
To configure the application, follow these steps:

1. **Create and Populate an `.ini` File**  
   - The `.ini` file must be located in the same directory as the `.exe` file.  
   - Multiple `.ini` files can exist; all of them will be processed.

2. **Obtain Your ApiKey**  
   - Copy the `ApiKey` from the eFaktura website under **Settings > API Management > Authentication Key**.  
   - Ensure the API status is set to **Active**.

3. **Set FolderPath**  
   - `FolderPath` specifies the location of `.xml` files that need to be sent.

### Example `.ini` File
```ini
[eFaktura]
ApiKey = "12345678-****-****-****-123456******"
FolderPath = "C:\Users\User1\Desktop\XML"
VatFolderPath = "C:\Users\User1\Desktop\XMLvat"
```

# How to Build and Run

## Running the Application
To run the application in development mode, use the following command:  
```bash
go run main.go
```

## Building demo executable
Run the following command to build demo executable the application:
```bash
go build -o efakturaDemo.exe
```
## Building executable
In the struct.go file, update the URL from demoefaktura to efaktura.

Run the following command to build executable the application:
```bash
go build -o efaktura.exe
```

# Notes

- **Invoices for CRF**  
  The XML file names for invoices intended for CRF must start with an uppercase letter **"B"**.

- **XML Structure Requirements**  
  Due to varying structures (especially for **final invoices** with specific extensions at the beginning), the XML must include the `<cbc:CustomizationID>` tag **before** `<cbc:ID>`.  
  This ensures the document number (**BrojDokumenta**) is read correctly.
