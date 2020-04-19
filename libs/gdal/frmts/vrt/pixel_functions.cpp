#include "gdal.h"

CPLErr Landsatv3CloudMasks(void **papoSources, int nSources, void *pData,
         int nXSize, int nYSize,
         GDALDataType eSrcType, GDALDataType eBufType,
         int nPixelSpace, int nLineSpace)
{
    int ii, iLine, iCol;
    GInt16 x0, x1, x2, pix_val;

    if (nSources != 3) return CE_Failure;

    for( iLine = 0; iLine < nYSize; iLine++ )
    {
        for( iCol = 0; iCol < nXSize; iCol++ )
        {
            ii = iLine * nXSize + iCol;
            x0 = SRCVAL(papoSources[0], eSrcType, ii);
            x1 = SRCVAL(papoSources[1], eSrcType, ii);
            x2 = SRCVAL(papoSources[2], eSrcType, ii);

            pix_val = x1 == 1 && x2 == 1 ? x0 : -999;

            GDALCopyWords(&pix_val, GDT_Int16, 0,
                        ((GByte *)pData) + nLineSpace * iLine + iCol * nPixelSpace,
                        eBufType, nPixelSpace, 1);
        }
    }

    return CE_None;
}

extern "C" void __attribute__((visibility("default"))) GDALRegister_GSKY_pixel_functions() {
  GDALAddDerivedBandPixelFunc("Landsatv3CloudMasks", Landsatv3CloudMasks);
}

