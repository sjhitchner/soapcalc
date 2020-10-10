from django.contrib import admin
from django.contrib.admin import DateFieldListFilter
from .models import (
    Additive,
    Fragrance,
    Supplier,
    Lipid,
    LipidInventory,
    Lye,
    LyeInventory,
    Recipe,
    RecipeAdditive,
    RecipeFragrance,
    RecipeLipid,
    RecipeBatch,
    RecipeBatchAdditive,
    RecipeBatchFragrance,
    RecipeBatchLipid,
    RecipeBatchLye,
)


class AdditiveAdmin(admin.ModelAdmin):
    pass

class FragranceAdmin(admin.ModelAdmin):
    pass

# class InventoryAdmin(admin.ModelAdmin):
#    pass

class SupplierAdmin(admin.ModelAdmin):
    pass

class LipidAdmin(admin.ModelAdmin):
    pass

class LipidInventoryAdmin(admin.ModelAdmin):
    pass

class LyeAdmin(admin.ModelAdmin):
    pass

class RecipeAdmin(admin.ModelAdmin):
    pass

class RecipeAdditiveAdmin(admin.ModelAdmin):
    pass

class RecipeFragranceAdmin(admin.ModelAdmin):
    pass

class RecipeLipidAdmin(admin.ModelAdmin):
    pass

class RecipeLyeAdmin(admin.ModelAdmin):
    pass

class RecipeBatchAdmin(admin.ModelAdmin):
    pass

class RecipeBatchAdditiveAdmin(admin.ModelAdmin):
    pass

class RecipeBatchFragranceAdmin(admin.ModelAdmin):
    pass

class RecipeBatchLipidAdmin(admin.ModelAdmin):
    pass

class RecipeBatchLyeAdmin(admin.ModelAdmin):
    pass


admin.site.register(Additive, AdditiveAdmin)
admin.site.register(Fragrance, FragranceAdmin)
# admin.site.register(Inventory, InventoryAdmin)
admin.site.register(Supplier, SupplierAdmin)
admin.site.register(Lipid, LipidAdmin)
admin.site.register(LipidInventory, LipidInventoryAdmin)
admin.site.register(Lye, LyeAdmin)
admin.site.register(Recipe, RecipeAdmin)
admin.site.register(RecipeAdditive, RecipeAdditiveAdmin)
admin.site.register(RecipeFragrance, RecipeFragranceAdmin)
admin.site.register(RecipeLipid, RecipeLipidAdmin)
admin.site.register(RecipeBatch, RecipeBatchAdmin)
admin.site.register(RecipeBatchAdditive, RecipeBatchAdditiveAdmin)
admin.site.register(RecipeBatchFragrance, RecipeBatchFragranceAdmin)
admin.site.register(RecipeBatchLipid, RecipeBatchLipidAdmin)
admin.site.register(RecipeBatchLye, RecipeBatchLyeAdmin)

