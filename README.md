# folderGen

application to quickly create a folder tree

You can create directory from creationFile and create templates scanning a directory

# Commands examples
folderGen create -f "B:\generadoresCarpeta\genCarpetaFac.json"

folderGen get -d "B:\Universidad\template" -o templateCarpetasFacultad -l B:/Universidad

-f path for creationFile
-d directory 
-o name output for template
-l location for save

# Formatt Creation File
```json
{
  "RootPaths": [
    "B:/Universidad/Cuarto Año/2do Cuatrimestre/Transmision De Datos/1-Teoria",
    "B:/Universidad/Cuarto Año/2do Cuatrimestre/Transmision De Datos/1-pichichi"
  ],
  "TemplatePath": "B:/generadoresCarpeta/Templates/templateCarpetasTeoria.json"
}
```
# Formatt template file
```json
{
 "Name": "",
 "SubFolders": [
  {
   "Name": "1-Teoria",
   "SubFolders": null
  },
  {
   "Name": "2-Practica",
   "SubFolders": [
    {
     "Name": "TrabajoPractico N1",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N2",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N3",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N4",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N5",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N6",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N7",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "TrabajoPractico N8",
     "SubFolders": [
      {
       "Name": "Solucion",
       "SubFolders": null
      }
     ]
    }
   ]
  },
  {
   "Name": "Evaluativos",
   "SubFolders": [
    {
     "Name": "Practica",
     "SubFolders": [
      {
       "Name": "Evaluativo 1",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 10",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 2",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 3",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 4",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 5",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 6",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 7",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 8",
       "SubFolders": null
      },
      {
       "Name": "Evaluativo 9",
       "SubFolders": null
      }
     ]
    },
    {
     "Name": "Teoria",
     "SubFolders": null
    }
   ]
  },
  {
   "Name": "Libros",
   "SubFolders": null
  },
  {
   "Name": "Parciales y Finales",
   "SubFolders": [
    {
     "Name": "Finales",
     "SubFolders": null
    },
    {
     "Name": "Parciales",
     "SubFolders": [
      {
       "Name": "1erParcial",
       "SubFolders": null
      },
      {
       "Name": "2doParcial",
       "SubFolders": null
      }
     ]
    }
   ]
  },
  {
   "Name": "Recursos",
   "SubFolders": null
  }
 ]
}
```
